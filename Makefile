linter-cache-clean:
	golangci-lint cache clean

.PHONY: lint
lint: linter-cache-clean
	golangci-lint run