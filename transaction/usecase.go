package transaction

import "context"

type UseCase interface {
	ListTransactions(ctx context.Context, userID int64) ([]Transaction, error)
	ListTransactionsByAccount(ctx context.Context, userID, accountID int64) ([]Transaction, error)
	CreateDepositTransaction(ctx context.Context, userID, accountID, amount int64) (*Transaction, error)
	CreateWithdrawTransaction(ctx context.Context, userID, accountID, amount int64) (*Transaction, error)
}
