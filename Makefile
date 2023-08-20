SHELL = /bin/bash

APP="ForkMatch"
GO_VERSION=1.21.0
GIT_HASH=$(shell git rev-parse --short HEAD)

# Compilation options:
# - CGO_ENABLED=0: Disable cgo
# - GOOS=linux: explicitly target Linux
# - GOARCH: explicitly target 64bit CPU
# - -trimpath: improve reproducibility by trimming the pwd from the binary
# - -ldflags: extra linker flags
#   - -s: omit the symbol table and debug information making the binary smaller
#   - -w: omit the DWARF symbol table making the binary smaller
#   - -extldflags 'static': extra linker flags: produce a statically linked binary
#   - -X main.Version=$VERSION: set the Version variable in the main package
# - -o dist/$(APP): output binary to dist/$(APP)
.PHONY: build
build: ## Build the binary
	@echo "Building binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
		go build \
		  -trimpath \
		  -ldflags "-s -w -extldflags '-static' -X main.Version=$(GIT_HASH)" \
		  -o dist/${APP} \
			./cmd/${APP} && \
		chmod +x dist/${APP}

.PHONY: test
test: ## Run unit tests
	@echo "Running tests..."
	@go test -v ./...

.PHONY: docker/build
docker/build: ## Build the docker image
	@echo "Building docker from git hash $(GIT_HASH)..."
	@docker build -t fork-match --build-arg GO_VERSION=$(GO_VERSION) --build-arg APP=$(APP) .

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
