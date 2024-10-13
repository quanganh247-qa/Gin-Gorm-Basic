package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/api/notes"
	"github.com/quanganh247-qa/gorm-project/app/api/users"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

func (server *Server) SetupRoutes(config util.Config) {
	routerDefault := gin.Default()

	routerDefault.Use(middleware.LoggingMiddleware())

	// api/v1
	v1 := routerDefault.Group(config.ApiPrefix)
	router := v1.Group("/")

	// Register routes
	routerGroup := middleware.RouterGroup{
		RouterDefault: router,
	}

	users.Routes(routerGroup)
	notes.Routes(routerGroup)

	server.Router = routerDefault

}
