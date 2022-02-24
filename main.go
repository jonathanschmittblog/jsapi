package main

import (
	"flag"

	"github.com/jonathanschmittblog/jsapi/routers"
	"github.com/jonathanschmittblog/jsapi/servers"
)

func main() {
	// Se estiver executando algum teste com "go test"
	if flag.Lookup("test.v") != nil {
		return
	}
	server, err := servers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor. Verifique as variáveis de ambiente.")
		return
	}
	routers.ApplyRoutes(server)
	server.Start()
}