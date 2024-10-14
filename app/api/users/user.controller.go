package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/util"
	"github.com/quanganh247-qa/gorm-project/app/util/token"
)

type UserServiceController interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	TestGetQuery(ctx *gin.Context)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": fmt.Sprintf("failed to bind request: %v", err),
		})
		return
	}
	user, err := c.service.CreateUserService(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("failed to create user: %v", err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"user": user,
	})
}

func (c *UserController) LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": fmt.Sprintf("failed to bind request: %v", err),
		})
		return
	}
	err := c.service.LoginUserService(ctx, req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": fmt.Sprintf("failed to login: %v", err),
		})
		return
	}
	acccessToken, _, err := token.TokenMaker.CreateToken(req.Username, util.Configs.AccessTokenDuration)
	ctx.JSON(200, gin.H{
		"token ": acccessToken,
	})
}

func (c *UserController) TestGetQuery(ctx *gin.Context) {
	pagination, err := util.GetPageInQuery(ctx.Request.URL.Query())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"pagination": pagination})
}
