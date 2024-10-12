package db

import (
	"context"
	"fmt"
)

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (s *Store) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	user := User{
		Username:     arg.Username,
		PasswordHash: arg.Password,
		Email:        arg.Email,
	}
	result := s.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return User{}, fmt.Errorf("error creating user: %w", result.Error)
	}
	return user, nil
}

func (s *Store) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var user User
	result := s.db.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return User{}, fmt.Errorf("error getting user: %w", result.Error)
	}
	return user, nil
}
