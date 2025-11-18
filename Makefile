.PHONY: help test test-go test-vue lint lint-go lint-vue build run coverage clean dev migrate docker-build docker-run

# Default target
.DEFAULT_GOAL := help

# Variables
BINARY_NAME=mahler
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-s -w -X main.version=$(VERSION)"

## help: Show this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## test: Run all tests
test: test-go test-vue

## test-go: Run Go tests with coverage
test-go:
	@echo "Running Go tests..."
	@go test -v -race -cover -coverprofile=coverage.out \
		./internal/app \
		./internal/api/... \
		./internal/project \
		./internal/service \
		./internal/resource/... \
		./internal/config \
		./internal

## test-vue: Run Vue tests with coverage
test-vue:
	@echo "Running Vue tests..."
	@cd web && npm test -- --run --coverage

## coverage: Check test coverage
coverage:
	@echo "Checking test coverage..."
	@./scripts/check-coverage.sh

## lint: Run all linters
lint: lint-go lint-vue

## lint-go: Run Go linters
lint-go:
	@echo "Running Go linters..."
	@golangci-lint run --timeout=5m

## lint-vue: Run Vue linters
lint-vue:
	@echo "Checking Vue formatting..."
	@cd web && npx prettier --check .

## fmt: Format all code
fmt:
	@echo "Formatting Go code..."
	@gofmt -s -w $(GO_FILES)
	@goimports -w $(GO_FILES)
	@echo "Formatting Vue code..."
	@cd web && npx prettier --write .

## build: Build the binary
build: build-frontend
	@echo "Building Go binary..."
	@go build $(LDFLAGS) -o $(BINARY_NAME) ./cmd
	@echo "Binary built: $(BINARY_NAME)"

## build-frontend: Build the Vue frontend
build-frontend:
	@echo "Building Vue frontend..."
	@cd web && npm run build

## run: Run the application
run:
	@echo "Running application..."
	@go run ./cmd

## dev: Run in development mode
dev:
	@echo "Starting development servers..."
	@$(MAKE) -j2 dev-backend dev-frontend

## dev-backend: Run backend in development mode
dev-backend:
	@echo "Starting backend..."
	@air || go run ./cmd

## dev-frontend: Run frontend in development mode
dev-frontend:
	@echo "Starting frontend..."
	@cd web && npm run dev

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(BINARY_NAME):$(VERSION) .
	@docker tag $(BINARY_NAME):$(VERSION) $(BINARY_NAME):latest

## docker-run: Run Docker container
docker-run:
	@echo "Running Docker container..."
	@docker run -p 8080:8080 $(BINARY_NAME):latest

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out
	@rm -f app_coverage.out
	@rm -rf web/dist
	@rm -rf web/coverage
	@echo "Clean complete"

## deps: Install dependencies
deps:
	@echo "Installing Go dependencies..."
	@go mod download
	@go mod verify
	@echo "Installing Vue dependencies..."
	@cd web && npm ci
	@echo "Dependencies installed"

## tidy: Tidy dependencies
tidy:
	@echo "Tidying Go dependencies..."
	@go mod tidy
	@echo "Tidying Vue dependencies..."
	@cd web && npm audit fix || true

## install-tools: Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "Tools installed"

## db-up: Start database (Docker Compose)
db-up:
	@echo "Starting database..."
	@docker-compose -f docker-compose.dev.yml up -d postgres
	@echo "Database started"

## db-down: Stop database
db-down:
	@echo "Stopping database..."
	@docker-compose -f docker-compose.dev.yml down
	@echo "Database stopped"

## version: Show version
version:
	@echo $(VERSION)
