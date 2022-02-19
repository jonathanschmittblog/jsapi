package routers

import (
	"jsapi/servers"
)

func ApplyAccountRoutes(server *servers.Server) {
	server.Router.POST("/pessoas", CreatePessoa)
	server.Router.PUT("/pessoas/:nome", UpdatePessoa)
	server.Router.GET("/pessoas/:nome", GetPessoa)
	server.Router.DELETE("/pessoas/:nome", DeletePessoa)
}