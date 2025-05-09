.PHONY: all build run clean fmt tidy help

# Makefile for Go project
all: build

build:
	mkdir -p bin
	go build -o bin/app ./cmd/app

run:
	go run cmd/app/main.go

clean:
	rm -rf bin
	go clean
	go mod tidy

fmt:
	go fmt ./...

tidy:
	go mod tidy

help:
	@echo "Makefile commands:"
	@echo "  all     - Build the project"
	@echo "  build   - Build the project"
	@echo "  run     - Run the project"
	@echo "  clean   - Clean up build artifacts"
	@echo "  fmt     - Format the code"
	@echo "  tidy    - Tidy up go modules"
	@echo "  help    - Show this help message"
