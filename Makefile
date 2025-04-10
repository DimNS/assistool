.EXPORT_ALL_VARIABLES:
GOBIN = $(shell pwd)/bin

.PHONY: deps
deps:
	@go mod tidy
	@go mod vendor

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test ./...

.PHONY: dev
dev:
	@./bin/wails dev

.PHONY: build
build: deps
	@./bin/wails build -platform windows/amd64 -o assistool.exe
	@./bin/wails build -platform linux/amd64 -o assistool_linux_amd64

.PHONY: tools
tools: deps
	@go install github.com/wailsapp/wails/v2/cmd/wails@latest
	@go install github.com/goreleaser/goreleaser@v1.21.2
