# Makefile for cross-platform build
GO = go

# Version information
VERSION ?= 0.1.1

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
