SHELL := /bin/bash
.DEFAULT_GOAL := help 

help: ## Show this help
	@echo Dependencies: go [upx]
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application in debug mode
	@go build -v

build_release: ## Build the application in production mode (requires UPX)
	@go build -ldflags="-s -w"
	@upx discord-png2fakegif.exe