package auth

import "context"

type UseCase interface {
	SignUp(ctx context.Context, email, password string) error
	SignIn(ctx context.Context, email, password string) (string, error)
}
