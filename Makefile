APP_NAME := orion
VERSION := v0.1.0
BUILD_DIR := dist

.PHONY: all build clean test release

all: build

build:
	@echo "Building..."
	go build -o $(APP_NAME) main.go

test:
	@echo "Running tests..."
	go test ./...

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f $(APP_NAME)

release: clean test
	@echo "Creating release..."
	mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_darwin_amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_darwin_arm64 main.go
	cd $(BUILD_DIR) && shasum -a 256 * > checksums.txt
	@echo "Release artifacts in $(BUILD_DIR)"
