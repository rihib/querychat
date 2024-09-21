clean:
	@echo "Cleaning..."
	@make clean-build
	@make clean-test

build: clean-build generate
	@echo "Building..."
	@go build -o ./bin/ ./...

run: build
	@echo "Running..."
	@./bin/$(APP_NAME)

clean-build:
	@echo "Cleaning build..."
	@rm -rf ./bin

test: clean-test generate
	@echo "Testing..."
	@go test -v ./...
	@make coverage

unit-test: clean-test generate
	@echo "Unit Testing..."
	@go test -v ./... -run Unit
	@make coverage

integration-test: clean-test generate
	@echo "Integration Testing..."
	@go test -v ./... -run Integration
	@make coverage

clean-test:
	@echo "Cleaning test cache..."
	@go clean -testcache
	@echo "Cleaning coverage..."
	@rm -rf test/coverage.out

generate:
	@echo "Generating..."
	@go generate ./...

coverage:
	@echo "Coverage..."
	@go test -coverprofile=test/coverage.out ./...
	@go tool cover -html=test/coverage.out

.PHONY: clean build run clean-build test unit-test integration-test clean-test generate coverage
