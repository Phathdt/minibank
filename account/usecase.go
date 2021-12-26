package account

import "context"

type UseCase interface {
	ListAccounts(ctx context.Context, userID int64) ([]Account, error)
	CreateAccount(ctx context.Context, userID, bankID int64, name string) (*Account, error)
}
