package transaction

import "context"

type UseCase interface {
	ListTransactions(ctx context.Context, userID int64) ([]Transaction, error)
	ListTransactionsByAccount(ctx context.Context, userId, accountID int64) ([]Transaction, error)
	//	CREATE transaction
}
