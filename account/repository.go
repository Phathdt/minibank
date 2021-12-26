package account

import "context"

type Repository interface {
	ListAccounts(ctx context.Context, userID int64) ([]Account, error)
}
