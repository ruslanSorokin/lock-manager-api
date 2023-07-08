.SILENT: tools.download
tools.download:
	@go mod download

.SILENT: tools.install
tools.install:
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.SILENT: generate.go.rest.echo
generate.go.rest.echo:
	./scripts/generate-go-server.sh echo

.SILENT: generate.go.rest.chi
generate.go.rest.chi:
	./scripts/generate-go-server.sh chi

.SILENT: generate.go.rest.gin
generate.go.rest.gin:
	./scripts/generate-go-server.sh gin

.SILENT: generate.go.rest.gorilla
generate.go.rest.gorilla:
	./scripts/generate-go-server.sh gorilla

.SILENT: generate.go.rest.fiber
generate.go.rest.fiber:
	./scripts/generate-go-server.sh fiber

.SILENT: generate.go.rest
generate.go.rest: generate.go.rest.echo generate.go.rest.chi generate.go.rest.gin generate.go.rest.gorilla generate.go.rest.fiber

GEN_GO_PREFIX := gen/grpc/go
SRC_PATH := src/grpc
SRC_PROTO := $(shell find $(SRC_PATH) -iname "*.proto")

.SILENT: generate.go.grpc
generate.go.grpc:
	@protoc \
	--proto_path=$(SRC_PATH) \
	\
	--go_out=$(GEN_GO_PREFIX) \
	--go-grpc_out=$(GEN_GO_PREFIX) \
	\
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	$(SRC_PROTO)


.SILENT: generate.go.
generate.go: generate.go.rest generate.go.grpc
