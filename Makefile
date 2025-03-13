# Makefile for cross-platform build
GO = go

# Version information
VERSION ?= 0.1.1

# Repository information
GITEE_OWNER ?= "oschina"
GITEE_REPO ?= "mcp-gitee"

# Flags
LDFLAGS = -ldflags "-X main.Version=$(VERSION)"
BUILD_FLAGS = -o bin/mcp-gitee $(LDFLAGS)

define show_usage_info
	@echo "\033[32m\n🤖🤖 Build Success 🤖🤖\033[0m"
	@echo "\033[32mExecutable path: $(shell pwd)/bin/mcp-gitee\033[0m"
	@echo "\033[33m\nUsage: ./bin/mcp-gitee [options]\033[0m"
	@echo "\033[33mAvailable options:\033[0m"
	@echo "\033[33m  --token=<token>       Gitee access token (or set GITEE_ACCESS_TOKEN env)\033[0m"
	@echo "\033[33m  --api-base=<url>      Gitee API base URL (or set GITEE_API_BASE env)\033[0m"
	@echo "\033[33m  --version             Show version information\033[0m"
	@echo "\033[33mExample: ./bin/mcp-gitee --token=your_access_token\033[0m"
	@echo "\033[33mExample with env: GITEE_ACCESS_TOKEN=your_token ./bin/mcp-gitee\033[0m"
endef

build:
	$(GO) build $(BUILD_FLAGS) -v main.go
	@echo "Build complete."
	$(call show_usage_info)

# Clean up generated binaries
clean:
	rm -f bin/mcp-gitee
	@echo "Clean up complete."

# 显示版本信息
version:
	@echo "Version: $(VERSION)"

# Clean up release directory
clean-release:
	rm -rf release
	@echo "Clean up release directory complete."

# Release artifacts for various system platforms
release: clean clean-release
	@mkdir -p release
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -v main.go
	@mv bin/mcp-gitee release/mcp-gitee-linux-amd64
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -v main.go
	@mv bin/mcp-gitee release/mcp-gitee-windows-amd64.exe
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -v main.go
	@mv bin/mcp-gitee release/mcp-gitee-darwin-amd64
	@echo "Building for macOS ARM..."
	GOOS=darwin GOARCH=arm64 $(GO) build $(BUILD_FLAGS) -v main.go
	@mv bin/mcp-gitee release/mcp-gitee-darwin-arm64
	@echo "Building for Linux ARM..."
	GOOS=linux GOARCH=arm $(GO) build $(BUILD_FLAGS) -v main.go
	@mv bin/mcp-gitee release/mcp-gitee-linux-arm
	@echo "Release complete. Artifacts are in the release directory."

# Upload artifacts to a specific release
upload-gitee-release:
	@echo "Uploading artifacts to gitee release..."
	@for file in release/*; do \
		curl -X POST \
			-H "Content-Type: multipart/form-data" \
			-F "access_token=$(GITEE_ACCESS_TOKEN)" \
			-F "owner=$(GITEE_OWNER)" \
			-F "repo=$(GITEE_REPO)" \
			-F "release_id=$(GITEE_RELEASE_ID)" \
			-F "file=@$$file" \
			https://gitee.com/api/v5/repos/$(GITEE_OWNER)/$(GITEE_REPO)/releases/$(GITEE_RELEASE_ID)/attach_files; \
	done
	@echo "Upload complete."
