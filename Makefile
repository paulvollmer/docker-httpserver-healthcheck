all: fmt lint test

fmt:
	@gofmt -s -w .

lint-prepare:
	@ curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1
lint:
	@ ./bin/golangci-lint run .
lint-fix:
	@ ./bin/golangci-lint run --fix .

test:
	@go test -v .
test-cov:
	@go test -v -cover .
test-cov-report:
	@go test -coverprofile=coverage.out .
	@go tool cover -html=coverage.out
