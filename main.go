package main

import (
	"flag"

	"github.com/jonathanschmittblog/jsapi/routers"
	"github.com/jonathanschmittblog/jsapi/servers"
)

func main() {
	server, err := servers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor. Verifique as variáveis de ambiente.")
		return
	}
	routers.ApplyRoutes(server)
	server.Start()
}