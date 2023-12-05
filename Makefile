BINARY_NAME=main
TERGET=main.go
BUILD_DIR=bin

.PHONY: check
check:
	@go fmt ./
	@if not exist $(BUILD_DIR) mkdir $(BUILD_DIR)

.PHONY: build-windows
build-windows: check
	@echo Building for windows ...
	@set GOOS=windows
	@set GOARCH=amd64
	@go build -o .\$(BUILD_DIR)\$(BINARY_NAME)-windows-amd64.exe $(TARGET)

.PHONY: build-linux
build-linux: check
	@echo Building for linux ...
	@set GOOS=linux
	@set GOARCH=amd64
	@go build -o .\$(BUILD_DIR)\$(BINARY_NAME)-linux-amd64 $(TARGET)

.PHONY: all
all: build-windows build-linux build-darwin

.PHONY: clean
clean:
	go clean
	@rmdir /s /q $(BUILD_DIR) 2>NUL


	
