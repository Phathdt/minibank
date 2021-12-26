package auth

import "errors"

var (
	ErrUserExist          = errors.New("User exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
)
