package users

import "github.com/quanganh247-qa/gorm-project/app/db"

type UserService struct {
	store db.Store
}

type UserAPI struct {
	controller UserServiceController
}

type UserController struct {
	service UserServiceInterface
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRepsonse struct {
	Token string `json:"token"`
}
