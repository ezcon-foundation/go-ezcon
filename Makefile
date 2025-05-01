
# Biến
BINARY_NAME=ezcon
BUILD_DIR=bin
SRC_DIR=cmd/ezcon
GO=go
GOFLAGS=-ldflags="-s -w"

# Mặc định: build
all: build

# Biên dịch chương trình
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go

# Chạy chương trình
run: build
	$(BUILD_DIR)/$(BINARY_NAME) --config ezcon.toml

# Cài đặt dependencies
deps:
	$(GO) mod tidy
	$(GO) mod download

# Xóa file build
clean:
	@rm -rf $(BUILD_DIR)

# Tạo file TOML mẫu
config:
	@echo 'node_id = "node1"' > ezcon.toml
	@echo 'private_key = "privkey1"' >> ezcon.toml
	@echo 'unl = ["node2:8081", "node3:8082", "node4:8083", "node5:8084"]' >> ezcon.toml
	@echo 'ledger_path = "./ledger.json"' >> ezcon.toml
	@echo 'rpc_port = "8080"' >> ezcon.toml
	@echo 'consensus_port = "9000"' >> ezcon.toml

# Kiểm tra mã nguồn
lint:
	@golangci-lint run

# Chạy unit test
test:
	$(GO) test ./... -v

# Giúp đỡ
help:
	@echo "Makefile for go-ezcon"
	@echo "Usage:"
	@echo "  make all       : Build the ezcon binary"
	@echo "  make build     : Build the ezcon binary"
	@echo "  make run       : Build and run the ezcon binary"
	@echo "  make deps      : Install dependencies"
	@echo "  make clean     : Remove build artifacts"
	@echo "  make config    : Generate sample ezcon.toml"
	@echo "  make lint      : Run linter"
	@echo "  make test      : Run unit tests"
	@echo "  make help      : Show this help"

.PHONY: all build run deps clean config lint test help
