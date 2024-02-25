.PHONY: all

PKG := "github.com/luyasr/gaia"

install: ## Install dependencies
	@go install github.com/swaggo/swag/cmd/swag@latest

swag: ## Generate swagger documentation
	@swag init

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help