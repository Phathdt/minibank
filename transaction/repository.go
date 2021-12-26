package transaction

import "context"

type Repository interface {
	ListTransactions(ctx context.Context, userID int64) ([]Transaction, error)
	ListTransactionsByAccount(ctx context.Context, userId, accountID int64) ([]Transaction, error)
	CreateTransaction(ctx context.Context, userID, accountID, amount int64, transactionType string) (*Transaction, error)
	GetAccount(ctx context.Context, userID, accountID int64) (*Account, error)
	UpdateAccount(ctx context.Context, accountID, balance int64) error
}
