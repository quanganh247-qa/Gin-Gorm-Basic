package middleware

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	RouterDefault *gin.RouterGroup
}

// func (rtGroup *RouterGroup) RouteAuth(router *gin.RouterGroup) gin.IRoutes {
// 	newRouter := router.Group("/")
// 	return newRouter.Use(AuthMiddleware(token.TokenMaker))
// }
