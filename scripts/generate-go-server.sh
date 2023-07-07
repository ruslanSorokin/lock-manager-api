for router in "$@"; do
  if [ "$router" = "echo" ]; then
    opt="server"
    package="ech"
  fi
  if [ "$router" = "chi" ]; then
    opt="chi-server"
    package="chi"
  fi
  if [ "$router" = "gin" ] ; then
    opt="$router"
    package="gin"
  fi
  if [ "$router" = "gorilla" ]; then
    opt="$router"
    package="grl"
  fi
  if [ "$router" = "fiber" ]; then
    opt="$router"
    package="fbr"
  fi
  rm -f "gen/rest/go/server/${router}/server.go"
  oapi-codegen -package "${package}api" -generate "types,${opt},spec" src/rest/swagger.yaml > "gen/rest/go/server/${router}/server.go"
done
