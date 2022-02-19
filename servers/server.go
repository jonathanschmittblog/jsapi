package servers

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port int
	Router *gin.Engine
}

func New() (*Server, error) {
	var err error
	server := &Server{}
	server.Port, err = strconv.Atoi(os.Getenv("JSAPI_PORT"))
	server.config()
	return server, err
}

func (server *Server) Start() {
	server.run()
}

func (server *Server) config() {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.Default()
}

func (server *Server) run() {
	println("Server is running at port " + strconv.Itoa(server.Port))
	err := server.Router.Run(":" + strconv.Itoa(server.Port))
	if err != nil {
		println(err)
	}
}