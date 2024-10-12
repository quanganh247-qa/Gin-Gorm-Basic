package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/util/token"
)

type RouterGroup struct {
	RouterDefault *gin.RouterGroup
}

func (rtGroup *RouterGroup) RouteAuth(router *gin.RouterGroup) gin.IRoutes {
	newRouter := router.Group("/")
	return newRouter.Use(AuthMiddleware(token.TokenMaker))
}
