APP_NAME=go-backend-api
MAIN_DIR=server
BIN_DIR=bin
PROTO_DIR=proto

# Colors for output
COLOR_RESET=\033[0m
COLOR_BOLD=\033[1m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m
COLOR_BLUE=\033[34m

. PHONY: run run-release build start clean proto deps test help

# Run in development mode with all Go files
run: 
	@echo "$(COLOR_GREEN)🚀 Running in DEV mode... $(COLOR_RESET)"
	@GIN_MODE=debug go run $(MAIN_DIR)/*.go

# Run in release mode
run-release:
	@echo "$(COLOR_GREEN)🚀 Running in RELEASE mode... $(COLOR_RESET)"
	@GIN_MODE=release go run $(MAIN_DIR)/*.go

# Build binary
build:
	@echo "$(COLOR_BLUE)🔨 Building binary...$(COLOR_RESET)"
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_DIR)/*.go
	@echo "$(COLOR_GREEN)✅ Build completed:  $(BIN_DIR)/$(APP_NAME)$(COLOR_RESET)"

# Build for production with optimizations
build-prod:
	@echo "$(COLOR_BLUE)🔨 Building production binary...$(COLOR_RESET)"
	@mkdir -p $(BIN_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $(BIN_DIR)/$(APP_NAME) $(MAIN_DIR)/*.go
	@echo "$(COLOR_GREEN)✅ Production build completed$(COLOR_RESET)"

# Start from binary
start:
	@echo "$(COLOR_GREEN)🚀 Starting from binary...$(COLOR_RESET)"
	@GIN_MODE=release ./$(BIN_DIR)/$(APP_NAME)

# Generate protobuf files
proto:
	@echo "$(COLOR_BLUE)📦 Generating protobuf files...$(COLOR_RESET)"
	@protoc --go_out=.  --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto
	@echo "$(COLOR_GREEN)✅ Protobuf generation completed$(COLOR_RESET)"

# Install dependencies
deps:
	@echo "$(COLOR_BLUE)📥 Installing dependencies...$(COLOR_RESET)"
	@go mod download
	@go mod tidy
	@echo "$(COLOR_GREEN)✅ Dependencies installed$(COLOR_RESET)"

# Install protobuf tools
install-proto:
	@echo "$(COLOR_BLUE)📥 Installing protobuf tools...$(COLOR_RESET)"
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "$(COLOR_GREEN)✅ Protobuf tools installed$(COLOR_RESET)"

# Run tests
test:
	@echo "$(COLOR_BLUE)🧪 Running tests...$(COLOR_RESET)"
	@go test -v ./... 

# Run tests with coverage
test-coverage:
	@echo "$(COLOR_BLUE)🧪 Running tests with coverage...$(COLOR_RESET)"
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(COLOR_GREEN)✅ Coverage report generated:  coverage.html$(COLOR_RESET)"

# Format code
fmt:
	@echo "$(COLOR_BLUE)📝 Formatting code...$(COLOR_RESET)"
	@go fmt ./...
	@echo "$(COLOR_GREEN)✅ Code formatted$(COLOR_RESET)"

# Lint code
lint:
	@echo "$(COLOR_BLUE)🔍 Linting code...$(COLOR_RESET)"
	@golangci-lint run ./... 

# Clean build artifacts
clean:
	@echo "$(COLOR_YELLOW)🧹 Cleaning...$(COLOR_RESET)"
	@rm -rf $(BIN_DIR)
	@rm -f coverage.out coverage.html
	@echo "$(COLOR_GREEN)✅ Cleaned$(COLOR_RESET)"

# Clean protobuf generated files
clean-proto:
	@echo "$(COLOR_YELLOW)🧹 Cleaning protobuf files...$(COLOR_RESET)"
	@rm -f $(PROTO_DIR)/*.pb.go
	@echo "$(COLOR_GREEN)✅ Protobuf files cleaned$(COLOR_RESET)"

# Full clean
clean-all: clean clean-proto
	@echo "$(COLOR_GREEN)✅ Full clean completed$(COLOR_RESET)"

# Development workflow - rebuild proto and run
dev:  proto run

# Production workflow - clean, build, and start
prod: clean-all proto build-prod

# Watch and auto-reload (requires air)
watch:
	@echo "$(COLOR_GREEN)👀 Starting live reload...$(COLOR_RESET)"
	@air

# Install air for live reload
install-air: 
	@echo "$(COLOR_BLUE)📥 Installing air... $(COLOR_RESET)"
	@go install github.com/cosmtrek/air@latest
	@echo "$(COLOR_GREEN)✅ Air installed$(COLOR_RESET)"

# Docker build
docker-build:
	@echo "$(COLOR_BLUE)🐳 Building Docker image...$(COLOR_RESET)"
	@docker build -t $(APP_NAME):latest .
	@echo "$(COLOR_GREEN)✅ Docker image built$(COLOR_RESET)"

# Docker run
docker-run:
	@echo "$(COLOR_GREEN)🐳 Running Docker container...$(COLOR_RESET)"
	@docker run -p 8080:8080 -p 50051:50051 $(APP_NAME):latest

# Show help
help:
	@echo "$(COLOR_BOLD)Available commands:$(COLOR_RESET)"
	@echo "  $(COLOR_GREEN)make run$(COLOR_RESET)              - Run in development mode"
	@echo "  $(COLOR_GREEN)make run-release$(COLOR_RESET)      - Run in release mode"
	@echo "  $(COLOR_GREEN)make build$(COLOR_RESET)            - Build binary"
	@echo "  $(COLOR_GREEN)make build-prod$(COLOR_RESET)       - Build optimized production binary"
	@echo "  $(COLOR_GREEN)make start$(COLOR_RESET)            - Start from binary"
	@echo "  $(COLOR_GREEN)make proto$(COLOR_RESET)            - Generate protobuf files"
	@echo "  $(COLOR_GREEN)make deps$(COLOR_RESET)             - Install dependencies"
	@echo "  $(COLOR_GREEN)make install-proto$(COLOR_RESET)    - Install protobuf tools"
	@echo "  $(COLOR_GREEN)make test$(COLOR_RESET)             - Run tests"
	@echo "  $(COLOR_GREEN)make test-coverage$(COLOR_RESET)    - Run tests with coverage"
	@echo "  $(COLOR_GREEN)make fmt$(COLOR_RESET)              - Format code"
	@echo "  $(COLOR_GREEN)make lint$(COLOR_RESET)             - Lint code"
	@echo "  $(COLOR_GREEN)make clean$(COLOR_RESET)            - Clean build artifacts"
	@echo "  $(COLOR_GREEN)make clean-proto$(COLOR_RESET)      - Clean protobuf files"
	@echo "  $(COLOR_GREEN)make clean-all$(COLOR_RESET)        - Full clean"
	@echo "  $(COLOR_GREEN)make dev$(COLOR_RESET)              - Generate proto and run"
	@echo "  $(COLOR_GREEN)make prod$(COLOR_RESET)             - Full production build"
	@echo "  $(COLOR_GREEN)make watch$(COLOR_RESET)            - Live reload (requires air)"
	@echo "  $(COLOR_GREEN)make install-air$(COLOR_RESET)      - Install air for live reload"
	@echo "  $(COLOR_GREEN)make docker-build$(COLOR_RESET)     - Build Docker image"
	@echo "  $(COLOR_GREEN)make docker-run$(COLOR_RESET)       - Run Docker container"
	@echo "  $(COLOR_GREEN)make help$(COLOR_RESET)             - Show this help"