
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
	@echo "Build Mitosisd"
	@go build -o build/mitosisd ./cmd/mitosisd
	@echo "Build Sidecar"
	@go build -o build/sidecar ./sidecar

release: test build
	@ignite chain build --output release --release --skip-proto

clean-local:
	@rm -rf ./test/localnet/*
	@rm -rf ./test/*.resp.json

setup-local: clean-local
	./test/local-setup.sh

run-local: setup-local
	@(./build/mitosisd start \
		--consensus.create_empty_blocks "false" \
		--p2p.pex "false" \
		--minimum-gas-prices "0.01umito" \
		--home "./test/localnet")

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
