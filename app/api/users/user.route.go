package users

import (
	"github.com/quanganh247-qa/gorm-project/app/api/middleware"
	"github.com/quanganh247-qa/gorm-project/app/db"
)

func Routes(routerGroup middleware.RouterGroup) {
	user := routerGroup.RouterDefault.Group("/user")

	userAPI := &UserAPI{
		&UserController{
			service: &UserService{
				store: db.StoreDB,
			},
		},
	}
	{
		user.POST("/create", userAPI.controller.CreateUser)
		user.POST("/login", userAPI.controller.LoginUser)
		user.GET("/test", userAPI.controller.TestGetQuery)
	}
}
