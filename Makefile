# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build binary
	@go build

test: unit integration ## Perform all unit and integration tests

unit: ## Perform unit tests
	@go test ./...

integration: build ## Perform integration tests
	@go test integration_test.go

generate: ## Generate go wordlists from diceware wordlists
	@for f in passphrase/*.txt; do rm -f "$${f%.txt}.go"; done
	@go generate

install: ## Install binary
	@go install
