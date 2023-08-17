APP="ForkMatch"
GO_VERSION=1.21.0

.PHONY: build
build: ## Build the binary
	@echo "Building binary..."
	@go build -o bin/$(APP) cmd/$(APP)/main.go

.PHONY: run
run: build ## Run the binary
	@echo "Running binary..."
	@./bin/$(APP)

.PHONY: test
test: ## Run unit tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: clean
clean: ## Clean the binary
	@echo "Cleaning..."
	@rm -rf bin/$(APP)

.PHONY: docker/build
docker/build: ## Build the docker image
	@echo "Building docker image..."
	@docker build -t fork-match --build-arg GO_VERSION=$(GO_VERSION) .

.PHONY: docker/run
docker/run: docker/build ## Run the docker image
	@echo "Running docker image..."
	@docker run -it --rm -p 80:80 fork-match

.PHONY: help
help: ## Display this help screen
	@echo "Usage: make [target] ..."
	@echo
	@echo "Targets:"
	@echo
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-30s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
