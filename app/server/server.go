package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

type Server struct {
	Router     *gin.Engine
	Connection *util.Connection
}

func NewServer(config util.Config) (*Server, error) {
	conn, err := util.Init(config)
	if err != nil {
		return nil, fmt.Errorf("can;t create new server")
	}

	server := &Server{
		Router: gin.Default(),
	}

	server.SetupRoutes(config)
	server.Connection = conn

	// routerDefault := gin.Default()

	// v1 := routerDefault.Group(config.ApiPrefix)
	// {
	// 	v1.GET("/test", func(ctx *gin.Context) {
	// 		ctx.JSON(200, gin.H{
	// 			"messange": "pong",
	// 		})
	// 	})

	// }

	// v1.Use(middleware.LoggingMiddleware())

	// server.Router = routerDefault

	return server, err

}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
