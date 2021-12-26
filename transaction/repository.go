package transaction

import "context"

type Repository interface {
	ListTransactions(ctx context.Context, userID int64) ([]Transaction, error)
	ListTransactionsByAccount(ctx context.Context, userId, accountID int64) ([]Transaction, error)
}
