
# `make` -> proto -> build
PHONY: build

clean:
	@rm -rf build
	@rm -rf release
	@rm -rf coverage.out

test:
	@go test -race -coverprofile=coverage.out.tmp -covermode=atomic ./...
	@cat coverage.out.tmp | grep -v '.pb.go' | grep -v '.pb.gw.go' > coverage.out
	@rm coverage.out.tmp

lint:
	@golangci-lint run

build: clean proto lint
	@ignite chain build --output build --skip-proto

release: test build
	@ignite chain build --output release --release --skip-proto

run-local: proto lint
	@ignite chain serve --skip-proto --quit-on-fail --verbose

run-local-clean: proto lint
	@ignite chain serve --skip-proto --quit-on-fail --verbose --reset-once

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
