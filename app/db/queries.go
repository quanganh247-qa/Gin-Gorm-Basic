package db

import "context"

type Queries interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
}
