package server

import (
	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/api/notes"
	"github.com/quanganh247-qa/gorm-project/app/api/users"
	"github.com/quanganh247-qa/gorm-project/app/util"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) SetupRoutes(config util.Config) {
	routerDefault := gin.Default()
	routerDefault.SetTrustedProxies(nil)
	routerDefault.Use(middleware.LoggingMiddleware())

	// api/v1
	v1 := routerDefault.Group(config.ApiPrefix)
	router := v1.Group("/")

	// Register routes
	routerGroup := middleware.RouterGroup{
		RouterDefault: router,
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	users.Routes(routerGroup)
	notes.Routes(routerGroup)

	server.Router = routerDefault

}
