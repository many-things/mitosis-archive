
# `make` -> proto -> build
PHONY: build

test:
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...

build: proto
	@ignite chain build --skip-proto

run-local: proto
	@ignite chain serve --skip-proto --quit-on-fail --verbose

proto: proto-fmt proto-go proto-openapi

proto-go:
	@echo "Generating protobuf bindings"
	@./scripts/protocgen.sh

proto-fmt:
	@echo "Formatting protobuf definitions"
	@buf format proto -w

proto-openapi:
	@echo "Generating OpenAPI document"
	@ignite generate openapi -y
