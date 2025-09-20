GO ?= go
TARGET ?= ./...
GOCACHE_DIR ?= $(PWD)/.gocache
GOLANGCI_LINT_CACHE_DIR ?= $(PWD)/.golangci-cache
GO_LINT_MODULE := github.com/skulidropek/devkitgo/cmd/lint@latest
GO_LINT_BIN ?= $(shell command -v go-lint 2>/dev/null)
GO_LINT_CMD := $(if $(GO_LINT_BIN),$(GO_LINT_BIN),$(GO) run $(GO_LINT_MODULE))

.PHONY: go-lint lint

go-lint lint:
	@mkdir -p $(GOCACHE_DIR) $(GOLANGCI_LINT_CACHE_DIR)
	GOCACHE=$(GOCACHE_DIR) GOLANGCI_LINT_CACHE=$(GOLANGCI_LINT_CACHE_DIR) $(GO_LINT_CMD) $(TARGET)
