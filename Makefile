# ==============================================================================
# Kavrok Makefile
# ==============================================================================

APP             := kavrok
MODULE          := github.com/mohammedimrankasab/kavrok
BUILD_DIR       := bin

GO              := go
GOFMT           := gofmt
GOLINT          := golangci-lint

VERSION         ?= dev
COMMIT          := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE            := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := \
	-s -w \
	-X main.version=$(VERSION) \
	-X main.commit=$(COMMIT) \
	-X main.buildDate=$(DATE)

.DEFAULT_GOAL := help

# ==============================================================================
# Help
# ==============================================================================

.PHONY: help
help: ## Show available commands
	@echo ""
	@echo "Kavrok Development Commands"
	@echo "==========================="
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

# ==============================================================================
# Development
# ==============================================================================

.PHONY: run
run: ## Run Kavrok
	$(GO) run ./cmd/kavrok

.PHONY: build
build: ## Build binary
	@mkdir -p $(BUILD_DIR)
	$(GO) build \
		-ldflags "$(LDFLAGS)" \
		-o $(BUILD_DIR)/$(APP) \
		./cmd/kavrok

.PHONY: install
install: ## Install binary
	$(GO) install ./cmd/kavrok

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)
	rm -f coverage.out
	rm -f coverage.html

# ==============================================================================
# Dependencies
# ==============================================================================

.PHONY: tidy
tidy: ## Tidy dependencies
	$(GO) mod tidy

.PHONY: download
download: ## Download dependencies
	$(GO) mod download

.PHONY: verify
verify: ## Verify dependencies
	$(GO) mod verify

# ==============================================================================
# Formatting
# ==============================================================================

.PHONY: fmt
fmt: ## Format source code
	$(GOFMT) -w .

.PHONY: vet
vet: ## Run go vet
	$(GO) vet ./...

# ==============================================================================
# Linting
# ==============================================================================

.PHONY: lint
lint: ## Run golangci-lint
	$(GO) tool $(GOLINT) run

# ==============================================================================
# Testing
# ==============================================================================

.PHONY: test
test: ## Run unit tests
	$(GO) test ./...

.PHONY: race
race: ## Run race detector
	$(GO) test -race ./...

.PHONY: coverage
coverage: ## Generate coverage report
	$(GO) test \
		-coverprofile=coverage.out \
		./...

	$(GO) tool cover \
		-html=coverage.out \
		-o coverage.html

	@echo ""
	@echo "Coverage report:"
	@echo "coverage.html"

# ==============================================================================
# Benchmarks
# ==============================================================================

.PHONY: benchmark
benchmark: ## Run benchmarks
	$(GO) test \
		-bench=. \
		-benchmem \
		./...

# ==============================================================================
# Quality
# ==============================================================================

.PHONY: check
check: fmt vet lint test ## Run all quality checks

# ==============================================================================
# Release
# ==============================================================================

.PHONY: release
release: clean check build ## Prepare release build

# ==============================================================================
# CI
# ==============================================================================

.PHONY: ci
ci: tidy check ## CI pipeline