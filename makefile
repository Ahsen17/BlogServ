# Go 编译器
GO := go
# 项目名称
APP_NAME := blogserv
# 源代码文件夹
#SRC_DIR := $(PWD)
SRC_DIR := $(CURDIR)
# 编译输出文件夹
BIN_DIR := $(SRC_DIR)/cmd

# 默认目标
default: build

# 编译项目
build:
	@echo "Building $(APP_NAME)..."
	$(GO) build -o $(BIN_DIR)/$(APP_NAME) $(SRC_DIR)/*.go
	@echo "Build complete."

# 运行项目
run: build
	@echo "Running $(APP_NAME)..."
	$(BIN_DIR)/$(APP_NAME)

# 清理编译产物
clean:
	@echo "Cleaning $(APP_NAME)..."
	rm -rf $(BIN_DIR)/*
	@echo "Clean complete."
