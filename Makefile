APP_NAME=go-backend-api
MAIN_FILE=cmd/server/main.go
BIN_DIR=bin

run:
	@echo "🚀 Running in DEV mode..."
	@GIN_MODE=debug go run $(MAIN_FILE)

run-release:
	@echo "🚀 Running in RELEASE mode..."
	@GIN_MODE=release go run $(MAIN_FILE)

build:
	@echo "🔨 Building binary..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_FILE)

start:
	@echo "🚀 Starting from binary..."
	@GIN_MODE=release ./$(BIN_DIR)/$(APP_NAME)

clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BIN_DIR)
