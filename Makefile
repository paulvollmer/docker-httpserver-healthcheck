all: fmt lint test

fmt:
	@gofmt -s -w .

lint:
	@golangci-lint run .

test:
	@go test -v .
test-cov:
	@go test -v -cover .
test-cov-report:
	@go test -coverprofile=coverage.out .
	@go tool cover -html=coverage.out
