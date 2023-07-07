package generate

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package "echapi" -generate "types,server,spec" -o "gen/rest/go/server/echo/server.go" "src/rest/swagger.yaml"

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package "chiapi" -generate "types,chi-server,spec" -o "gen/rest/go/server/chi/server.go" "src/rest/swagger.yaml"

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package "ginapi" -generate "types,gin,spec" -o "gen/rest/go/server/gin/server.go" "src/rest/swagger.yaml"

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package "grlapi" -generate "types,gorilla,spec" -o "gen/rest/go/server/gorilla/server.go" "src/rest/swagger.yaml"

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest -package "fbrapi" -generate "types,fiber,spec" -o "gen/rest/go/server/fiber/server.go" "src/rest/swagger.yaml"
