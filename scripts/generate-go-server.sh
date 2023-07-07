for router in "$@"; do
  if [ "$router" = "echo" ]; then
    opt="server"
  fi
  if [ "$router" = "chi" ]; then
    opt="chi-server"
  fi
  if [ "$router" = "gin" ] || [ "$router" = "gorilla" ] || [ "$router" = "fiber" ]; then
    opt="$router"
  fi
  rm -f "gen/rest/go/server/${router}/server.go"
  oapi-codegen -package "${router}api" -generate "types,${opt},spec" src/rest/swagger.yaml > "gen/rest/go/server/${router}/server.go"
done
