.PHONY: help setup run test build docker-build docker-run docker-kill lint update-dependencies

help: ## Show help.
	@printf "A set of development commands.\n"
	@printf "\nUsage:\n"
	@printf "\t make \033[36m<commands>\033[0m\n"
	@printf "\nThe Commands are:\n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup server.
	@go mod download

run: ## Run local server.
	@go run main.go

test: ## Run test.
	@go test -v ./... -covermode=atomic -count=1

build: ## Build server.
	@go build -v -o main .

build-consumer: ## Build consumer client.
	@go build -v -o clientConsumer ./consumer/.

run-consumer: ## Run consumer client.
	@go run ./consumer/consumer.go

docker-build: ## Build container's Docker.
	@docker build -t server .

docker-run: ## Run container's Docker.
	@docker run --name CatMQ -p 23023:23023 -it server

docker-kill: ## Kill container's Docker.
	@docker kill CatMQ

lint: ## Run lint.
	golangci-lint run ./... --enable-all

update-dependencies: ## Update all dependencies.
	@go get -u
