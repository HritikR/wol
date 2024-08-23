# Source file
SRC=main.go

# Build Directory
BUILD_DIR=./build

# The default cmd
all: build

# Build
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build complete."

# Run
run: build
	@echo "Running the WOL..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Clean the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."

# Phony targets to avoid conflicts with file names
.PHONY: all build run install clean
