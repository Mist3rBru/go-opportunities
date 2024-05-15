.PHONY: default run build test docs clean
# Variables
APP_NAME=gopportunities

# Tasks
default: run

run:
	@go run cmd/api/main.go
build:
	@go build -a -o $(APP_NAME) cmd/api/main.go
test:
	@go test ./ ...
clean:
	@rm -f $(APP_NAME)
format:
	@gofmt -w .