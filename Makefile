# Makefile for challenge-fravega

# Variables
APP_NAME := app
CMD_DIR := ./cmd/server
DOCKER_IMAGE := challenge-fravega

# Go related variables
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

.PHONY: all build clean test run lint fmt docker-build docker-run docker-compose-up docker-compose-down help

all: clean build ## Build the application

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o $(GOBIN)/$(APP_NAME) $(CMD_DIR)

clean: ## Remove previous build
	@echo "Cleaning..."
	@rm -f $(GOBIN)/$(APP_NAME)
	@go clean

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

run: build ## Run the application
	@echo "Running $(APP_NAME)..."
	@$(GOBIN)/$(APP_NAME)

lint: ## Run linter
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed"; \
		go vet ./...; \
	fi

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build ## Run Docker container
	@echo "Running Docker container..."
	@docker run -p 8080:8080 $(DOCKER_IMAGE)

docker-compose-up: ## Start all services with docker-compose
	@echo "Starting services with docker-compose..."
	@docker-compose up -d

docker-compose-down: ## Stop all services with docker-compose
	@echo "Stopping services with docker-compose..."
	@docker-compose down

mock-up: ## Start only mock service with docker-compose
	@echo "Starting mock service..."
	@docker-compose up -d mmock

db-up: ## Start database services with docker-compose
	@echo "Starting database services..."
	@docker-compose up -d h2 redis

db-migrate: ## Run database migrations
	@echo "Running database migrations..."
	@docker-compose up -d dbmate

# Purchase Order Mock Targets
po-mock-up: ## Start the purchase order mock server
	@echo "Starting purchase order mock server..."
	@cd test/mocks && docker-compose up -d purchaseorder-mock

po-mock-down: ## Stop the purchase order mock server
	@echo "Stopping purchase order mock server..."
	@cd test/mocks && docker-compose down

po-mock-logs: ## View the purchase order mock server logs
	@echo "Viewing purchase order mock server logs..."
	@cd test/mocks && docker-compose logs -f purchaseorder-mock

po-example: ## Run the purchase order mock example
	@echo "Running purchase order mock example..."
	@cd test/mocks/purchaseorder && go run example_usage.go

tidy: ## Tidy and vendor Go modules
	@echo "Tidying Go modules..."
	@go mod tidy

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.DEFAULT_GOAL := help 