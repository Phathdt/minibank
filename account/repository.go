package account

import "context"

type Repository interface {
	ListAccounts(ctx context.Context, userID int64) ([]Account, error)
	GetAccount(ctx context.Context, accountID int64) (*Account, error)
	UpdateAccount(ctx context.Context, accountID, balance int64) error
}
