all: fmt lint test

fmt:
	@gofmt -s -w .

lint:
	@golangci-lint run .

test:
	@go test -v .
