package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quanganh247-qa/gorm-project/app/db"
	"github.com/quanganh247-qa/gorm-project/app/util"
)

type UserServiceInterface interface {
	CreateUserService(ctx *gin.Context, req CreateUserRequest) (*db.User, error)
	LoginUserService(ctx *gin.Context, req LoginRequest) error
}

func (s *UserService) CreateUserService(ctx *gin.Context, req CreateUserRequest) (*db.User, error) {
	hashedPassword, err := util.HassPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Email:    req.Email,
		Username: req.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &user, nil
}

func (s *UserService) LoginUserService(ctx *gin.Context, req LoginRequest) error {
	user, err := s.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if err := util.VerifyPassword(req.Password, user.PasswordHash); err != nil {
		return fmt.Errorf("failed to verify password: %w", err)
	}

	return nil
}
