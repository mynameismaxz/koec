SHELL := /bin/bash

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the application
	$(call print-target)
	@go run main.go

.PHONY: build
build: ## Build the binary file application
	$(call print-target)
	@go build -x -o bin/koec main.go

.PHONY: clean
clean: ## Clean the binary file application
	$(call print-target)
	@rm -rf bin

.PHONY: github-validate
github-validate: ## Run act to validate github actions
	$(call print-target)
	@act -n

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef