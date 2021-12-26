package transaction

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountBalance  = errors.New("balance not enough")
)
