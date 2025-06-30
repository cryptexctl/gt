.PHONY: build clean test install

BINARY_NAME=gt
BUILD_DIR=bin

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go

clean:
	rm -rf $(BUILD_DIR)

test:
	go test -v ./...

install: build
	sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

dev: build
	./$(BUILD_DIR)/$(BINARY_NAME)

fmt:
	go fmt ./...

vet:
	go vet ./...

mod-tidy:
	go mod tidy

help:
	@echo "Доступные команды:"
	@echo "  build      - сборка бинарника"
	@echo "  build-all  - сборка для всех платформ"
	@echo "  clean      - очистка"
	@echo "  test       - запуск тестов"
	@echo "  install    - установка в систему"
	@echo "  dev        - сборка и запуск"
	@echo "  fmt        - форматирование кода"
	@echo "  vet        - проверка кода"
	@echo "  mod-tidy   - обновление зависимостей" 