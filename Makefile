# Makefile for managing project tasks

.PHONY: setup build run lint

# Set up the environment by downloading dependencies
setup:
	go mod tidy

# Build the project
build:
	go build -o ./bin/app ./cmd

# Run the application
run: build
	./bin/app

# Run the linter
lint:
	golangci-lint run

# Run the tests
test:
	go test -v ./...

# Run the tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

# Clean up the build artifacts
clean:
	rm -rf ./bin/*