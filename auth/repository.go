package auth

import (
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, password string) error
	GetUser(ctx context.Context, id int) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
