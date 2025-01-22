PROTOC_VERSION ?= 24.0
.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: proto-clean
proto-clean: ## Clean generated proto.
	@rm -rf pb/v1

.PHONY: proto-compile
proto-compile: ## Compile message protobuf and gRPC service files.
	PLATFORM=$(shell uname -m) PROTOC_VERSION=$(PROTOC_VERSION) docker-compose -f docker/docker-compose.yaml up --build -d

.PHONY: docker-config
docker-config: ## Dump docker-compose configuration.
	PLATFORM=$(shell uname -m) PROTOC_VERSION=$(PROTOC_VERSION) docker-compose -f docker/docker-compose.yaml config