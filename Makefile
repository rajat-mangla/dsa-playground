# Go DSA Playground Makefile

APP_NAME=dsa-playground

.PHONY: run test tidy fmt lint clean

## Run main.go
run:
	@echo "▶️  Running main.go..."
	go run main.go

## Run all tests
test:
	@echo "🧪 Running tests..."
	go test ./... -v

## Format code
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...

## Tidy up go.mod and go.sum
tidy:
	@echo "🧹 Tidying module dependencies..."
	go mod tidy

## Lint code (if golangci-lint installed)
lint:
	@echo "🔍 Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "⚠️ golangci-lint not installed. Install with: brew install golangci-lint"; \
	fi

## Clean build cache
clean:
	@echo "🧼 Cleaning build cache..."
	go clean -cache -testcache -modcache

