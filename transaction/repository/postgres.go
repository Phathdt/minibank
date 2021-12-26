package repository

import (
	"context"
	"database/sql"

	"minibank/db/postgresql"
	"minibank/transaction"
)

type TransactionRepo struct {
	q *postgresql.Queries
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{q: postgresql.New(db)}
}

func (r TransactionRepo) ListTransactions(ctx context.Context, userID int64) ([]transaction.Transaction, error) {
	ts, err := r.q.ListTransactions(ctx, userID)
	if err != nil {
		return nil, err
	}

	transactions := make([]transaction.Transaction, len(ts))

	for i, t := range ts {
		transactions[i] = t.MapToEntity()
	}

	return transactions, nil
}

func (r TransactionRepo) ListTransactionsByAccount(ctx context.Context, userId, accountID int64) ([]transaction.Transaction, error) {
	ts, err := r.q.ListTransactionsByAccount(ctx, postgresql.ListTransactionsByAccountParams{
		ID:   userId,
		ID_2: accountID,
	})
	if err != nil {
		return nil, err
	}

	transactions := make([]transaction.Transaction, len(ts))

	for i, t := range ts {
		transactions[i] = t.MapToEntity()
	}

	return transactions, nil
}
