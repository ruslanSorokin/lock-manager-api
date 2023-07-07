.SILENT: tools.download
tools.download:
	@go mod download

.SILENT: tools.install
tools.install:
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.SILENT: go.generate.echo
go.generate.echo:
	./scripts/generate-go-server.sh echo

.SILENT: go.generate.chi
go.generate.chi:
	./scripts/generate-go-server.sh chi

.SILENT: go.generate.gin
go.generate.gin:
	./scripts/generate-go-server.sh gin

.SILENT: go.generate.gorilla
go.generate.gorilla:
	./scripts/generate-go-server.sh gorilla

.SILENT: go.generate.fiber
go.generate.fiber:
	./scripts/generate-go-server.sh fiber

.SILENT: go.generate.client
go.generate.client:
	@oapi-codegen -package "clientapi" -generate "types,client,spec" src/rest/swagger.yaml > "gen/rest/go/client/client.go"

.SILENT: go.generate
go.generate: go.generate.echo go.generate.chi go.generate.gin go.generate.gorilla go.generate.fiber go.generate.client
