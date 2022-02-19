package main

import (
	"jsapi/routers"
	"jsapi/servers"
)

func main() {
	server, err := servers.New()
	if err != nil {
		println("Não foi possível iniciar o servidor. Verifique as variáveis de ambiente.")
		return
	}
	routers.ApplyAccountRoutes(server)
	server.Start()
}