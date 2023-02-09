PKG := "github.com/psyb0t/gopenai"
PKG_LIST := $(shell go list $(PKG)/...)

dep: ## Get the dependencies + remove unused ones
	@go mod tidy
	@go mod download

lint: ## Lint Golang files
	@golangci-lint run --timeout=30m0s

test: ## Run tests
	@go test -race -v $(PKG_LIST)

test-coverage: ## Run tests with coverage
	@go test -race -coverprofile coverage.txt -covermode=atomic ${PKG_LIST}

test-coverage-tool: test-coverage ## Run test coverage followed by the cover tool
	@go tool cover -func=coverage.txt
	@go tool cover -html=coverage.txt

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
