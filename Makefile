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

release:
	@echo "Creating release..."
	goreleaser release --clean

snapshot:
	@echo "Creating manual snapshot release..."
	goreleaser release --snapshot --clean
