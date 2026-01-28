.PHONY: help build run test clean docker-build docker-up docker-down dev-up dev-down install-air

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	go build -o bin/main ./cmd/api

run: ## Run the application locally
	go run ./cmd/api/main.go

test: ## Run tests
	go test -v -race -coverprofile=coverage.out ./...

coverage: test ## Generate coverage report
	go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	rm -rf bin/ tmp/ coverage.out coverage.html

lint: ## Run linters
	go fmt ./...
	go vet ./...

docker-build: ## Build Docker image
	docker build -t golang-app:latest .

docker-up: ## Start production docker-compose
	docker-compose up -d

docker-down: ## Stop production docker-compose
	docker-compose down

docker-logs: ## Show docker-compose logs
	docker-compose logs -f

dev-up: ## Start development environment
	docker-compose -f docker-compose.dev.yml up --build

dev-down: ## Stop development environment
	docker-compose -f docker-compose.dev.yml down

dev-logs: ## Show development logs
	docker-compose -f docker-compose.dev.yml logs -f

install-deps: ## Install Go dependencies
	go mod download
	go mod tidy

install-air: ## Install air for hot reload
	go install github.com/air-verse/air@latest

migrate: ## Run database migrations
	@echo "Running migrations..."
	go run ./cmd/api/main.go

deps: ## Download dependencies
	go mod download
	go mod verify
