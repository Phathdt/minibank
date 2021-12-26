package account

import "context"

type UseCase interface {
	ListAccounts(ctx context.Context, userID int64) ([]Account, error)
}
