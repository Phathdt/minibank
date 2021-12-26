package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"minibank/auth"

	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

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

func (au *AuthUseCase) SignUp(ctx context.Context, username, password string) error {
	user, _ := au.userRepo.GetUserByUsername(ctx, username)
	if user != nil {
		return auth.ErrUserExist
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))

	return au.userRepo.CreateUser(ctx, username, fmt.Sprintf("%x", pwd.Sum(nil)))
}

func (au *AuthUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	user, err := au.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	if password != user.Password {
		return "", auth.ErrPasswordNotMatch
	}

	claims := AuthClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(au.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(au.signingKey)
}
