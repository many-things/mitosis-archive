
proto: proto-fmt proto-go proto-openapi

proto-go:
	./scripts/protocgen.sh

proto-fmt:
	buf format proto -w

proto-openapi:
	ignite generate openapi -y
