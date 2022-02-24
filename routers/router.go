package routers

import (
	"github.com/jonathanschmittblog/jsapi/servers"
)

func ApplyRoutes(server *servers.Server) {
	server.Router.POST("/pessoas", CreatePessoa)
	server.Router.PUT("/pessoas/:nome", UpdatePessoa)
	server.Router.GET("/pessoas/:nome", GetPessoa)
	server.Router.DELETE("/pessoas/:nome", DeletePessoa)
}