package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"minibank/auth"
)

type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(userRepo auth.UserRepository, hashSalt string, signingKey []byte, expireDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signingKey:     signingKey,
		expireDuration: expireDuration,
	}
}

func (au *AuthUseCase) SignUp(ctx context.Context, email, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))

	return au.userRepo.CreateUser(ctx, email, fmt.Sprintf("%x", pwd.Sum(nil)))
}

func (a *AuthUseCase) SignIn(ctx context.Context, email, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}
