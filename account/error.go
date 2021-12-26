package account

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrNotYourAccount  = errors.New("not your account")
)
