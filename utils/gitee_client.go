package utils

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"

	"gitee.com/oschina/mcp-gitee/operations/types"

	"github.com/mark3labs/mcp-go/mcp"
)

const (
	DefaultApiBase = "https://gitee.com/api/v5"
)

var (
	giteeAccessToken string
	apiBase          string
)

func SetGiteeAccessToken(token string) {
	giteeAccessToken = token
}

func SetApiBase(url string) {
	apiBase = url
}

func GetGiteeAccessToken() string {
	if giteeAccessToken != "" {
		return giteeAccessToken
	}
	return os.Getenv("GITEE_ACCESS_TOKEN")
}

func GetApiBase() string {
	if apiBase != "" {
		return apiBase
	}
	if envApiBase := os.Getenv("GITEE_API_BASE"); envApiBase != "" {
		return envApiBase
	}
	return DefaultApiBase
}

type GiteeClient struct {
	Url       string
	Method    string
	Payload   interface{}
	Headers   map[string]string
	Response  *http.Response
	parsedUrl *url.URL
	Query     map[string]string
	SkipAuth  bool
	Ctx       context.Context
}

type Option func(client *GiteeClient)

type ErrMsgV5 struct {
	Message string `json:"message"`
}

func NewGiteeClient(method, urlString string, opts ...Option) *GiteeClient {
	urlString = GetApiBase() + urlString
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	client := &GiteeClient{
		Method:    method,
		Url:       parsedUrl.String(),
		parsedUrl: parsedUrl,
		Ctx:       context.Background(),
	}

	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithQuery(query map[string]interface{}) Option {
	return func(client *GiteeClient) {
		parsedQuery := make(map[string]string)
		if query != nil {
			queryParams := client.parsedUrl.Query()
			for k, v := range query {
				parsedValue := ""
				switch v.(type) {
				case string:
					parsedValue = v.(string)
				case int:
					parsedValue = strconv.Itoa(v.(int))
				case float32, float64:
					parsedValue = fmt.Sprintf("%.f", v)
				case bool:
					parsedValue = strconv.FormatBool(v.(bool))
				}
				if parsedValue != "" {
					queryParams.Set(k, parsedValue)
					parsedQuery[k] = parsedValue
				}
			}
			client.parsedUrl.RawQuery = queryParams.Encode()
		}
		client.Url = client.parsedUrl.String()
		client.Query = parsedQuery
	}
}

func WithPayload(payload interface{}) Option {
	return func(client *GiteeClient) {
		client.Payload = payload
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(client *GiteeClient) {
		client.Headers = headers
	}
}

func WithSkipAuth() Option {
	return func(client *GiteeClient) {
		client.SkipAuth = true
	}
}

func (g *GiteeClient) SetHeaders(headers map[string]string) *GiteeClient {
	g.Headers = headers
	return g
}

func (g *GiteeClient) Do() (*GiteeClient, error) {
	g.Response = nil
	_payload, _ := json.Marshal(g.Payload)
	req, err := http.NewRequest(g.Method, g.Url, bytes.NewReader(_payload))
	if err != nil {
		return nil, NewInternalError(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "mcp-gitee "+Version+" Go/"+runtime.GOOS+"/"+runtime.GOARCH+"/"+runtime.Version())

	accessToken := ""
	if g.Ctx != nil {
		if v := g.Ctx.Value("access_token"); v != nil {
			if s, ok := v.(string); ok && s != "" {
				accessToken = s
			}
		}
	}
	if accessToken == "" {
		accessToken = GetGiteeAccessToken()
	}
	if accessToken == "" && !g.SkipAuth {
		return nil, NewAuthError()
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return g, NewNetworkError(err)
	}

	g.Response = resp

	if !g.IsSuccess() {
		body, _ := ioutil.ReadAll(resp.Body)
		return g, NewAPIError(resp.StatusCode, body)
	}

	return g, nil
}

func (g *GiteeClient) IsSuccess() bool {
	if g.Response == nil {
		return false
	}

	successMap := map[int]struct{}{
		http.StatusOK:          struct{}{},
		http.StatusCreated:     struct{}{},
		http.StatusNoContent:   struct{}{},
		http.StatusFound:       struct{}{},
		http.StatusNotModified: struct{}{},
	}

	if _, ok := successMap[g.Response.StatusCode]; ok {
		return true
	}
	return false
}

func (g *GiteeClient) IsFail() bool {
	return !g.IsSuccess()
}

func (g *GiteeClient) GetRespBody() ([]byte, error) {
	return ioutil.ReadAll(g.Response.Body)
}

func (g *GiteeClient) HandleMCPResult(object any) (*mcp.CallToolResult, error) {
	_, err := g.Do()
	if err != nil {
		switch {
		case IsAuthError(err):
			return mcp.NewToolResultError("Authentication failed: Please check your Gitee access token"), err
		case IsNetworkError(err):
			return mcp.NewToolResultError("Network error: Unable to connect to Gitee API"), err
		case IsAPIError(err):
			giteeErr := err.(*GiteeError)
			return mcp.NewToolResultError(fmt.Sprintf("API error (%d): %s", giteeErr.Code, giteeErr.Details)), err
		default:
			return mcp.NewToolResultError(err.Error()), err
		}
	}

	if object == nil {
		return mcp.NewToolResultText("Operation completed successfully"), nil
	}

	body, err := g.GetRespBody()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to read response body: %s", err.Error())),
			NewInternalError(errors.New(err.Error()))
	}

	if err = json.Unmarshal(body, object); err != nil {
		errorMessage := fmt.Sprintf("Failed to parse response: %v", err)
		return mcp.NewToolResultError(errorMessage), NewInternalError(errors.New(errorMessage))
	}

	switch v := object.(type) {
	case *[]types.FileContent:
		for i := range *v {
			content, err := base64.StdEncoding.DecodeString((*v)[i].Content)
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to decode base64 content: %s", err.Error())),
					NewInternalError(err)
			}
			(*v)[i].Content = string(content)
		}
		object = v
	case *types.FileContent:
		content, err := base64.StdEncoding.DecodeString(v.Content)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to decode base64 content: %s", err.Error())),
				NewInternalError(err)
		}
		v.Content = string(content)
		object = v
	}

	result, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %s", err.Error())),
			NewInternalError(err)
	}

	return mcp.NewToolResultText(string(result)), nil
}

func WithContext(ctx context.Context) Option {
	return func(client *GiteeClient) {
		if ctx != nil {
			client.Ctx = ctx
		}
	}
}
