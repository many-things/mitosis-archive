
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

setup-local:
	@rm -rf ./test/localnet/*

	@(./build/mitosisd init localnet \
		--chain-id 'mito-local-1' \
		--staking-bond-denom 'umito' \
		--home "./test/localnet")

	@sed -i '' 's/stake/umito/g' ./test/localnet/config/genesis.json

	@((echo "maple often cargo polar eager jaguar eight inflict once nest nice swamp weasel address swift physical valid culture cheese trumpet find dinosaur curve tray"; echo "mitomito"; echo "mitomito") \
		| ./build/mitosisd keys add validator \
		--recover \
		--home "./test/localnet" \
		--keyring-backend "file")

	@(echo "mitomito" \
		| ./build/mitosisd add-genesis-account \
		validator 2000000000000umito \
		--home "./test/localnet" \
		--keyring-backend "file")

	@(echo "mitomito" \
		| ./build/mitosisd gentx validator 1000000000000umito \
		--chain-id 'mito-local-1' \
		--keyring-backend "file" \
		--home "./test/localnet")

	@(./build/mitosisd collect-gentxs --home "./test/localnet")

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
