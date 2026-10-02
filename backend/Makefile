# Variables
APP_NAME=gochat
MAIN_PATH=cmd/main.go
BINARY_PATH=bin/$(APP_NAME)

# Colors
GREEN=\033[0;32m
NC=\033[0m # No Color

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' Makefile | sed 's/## /  /'

# ---------- GO UTILS ----------
## GO: ---------- :GO

## tidy: Tidy go modules
tidy:
	@echo "$(GREEN)Tidying modules...$(NC)"
	@go mod tidy

## fmt: Format code
fmt:
	@echo "$(GREEN)Formatting code...$(NC)"
	@go fmt ./...
	@goimports -local github.com/Din18m/GOnnect -w .

## lint: Run linter
lint:
	@echo "$(GREEN)Running linter...$(NC)"
	@golangci-lint run ./...

## lint-fix: Run linter and fix issues
lint-fix:
	@echo "$(GREEN)Running linter with auto-fix...$(NC)"
	@golangci-lint run --fix ./...

# ---------- TESTS ----------
## TESTS: ---------- :TESTS

## test: Run tests
test:
	@echo "$(GREEN)Running tests...$(NC)"
	@go test -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)Coverage report: coverage.html$(NC)"


## clean: Clean build artifacts
clean:
	@echo "$(GREEN)Cleaning...$(NC)"
	@rm -rf bin/ tmp/
	@rm -f coverage.out coverage.html
	@echo "$(GREEN)Cleaned$(NC)"

# ---------- LOCAL DEVELOPMENT ----------
## LOCAL: ---------- :LOCAL

## dev: Setting dev environment and launch air for local development
dev:
	set -a && . ./.env && set +a && air

## all-up: Launching all services with build in background
all-up:
	docker compose up -d --build

## all-down: Stopping all services
all-down:
	docker compose down

## setup: Install development dependencies
setup:
	@echo "$(GREEN)Installing development dependencies...$(NC)"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/air-verse/air@latest
	@echo "$(GREEN)Dependencies installed$(NC)"
