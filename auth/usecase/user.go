package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"minibank/auth"
)

type AuthClaims struct {
	jwt.StandardClaims
	ID    int32  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
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

func (au *AuthUseCase) SignUp(ctx context.Context, email, password string) error {
	user, _ := au.userRepo.GetUserByEmail(ctx, email)
	if user != nil {
		return auth.ErrUserExist
	}

	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(au.hashSalt))

	return au.userRepo.CreateUser(ctx, email, fmt.Sprintf("%x", pwd.Sum(nil)))
}

func (au *AuthUseCase) SignIn(ctx context.Context, email, password string) (string, error) {
	user, err := au.userRepo.GetUserByEmail(ctx, email)
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
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(au.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(au.signingKey)
}
