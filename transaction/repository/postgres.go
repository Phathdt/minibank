package repository

import (
	"context"
	"database/sql"

	"minibank/db/postgresql"
	"minibank/transaction"
)

type Repo struct {
	q *postgresql.Queries
}

func NewTransactionRepo(db *sql.DB) *Repo {
	return &Repo{q: postgresql.New(db)}
}

func (r Repo) CreateTransaction(ctx context.Context, userID, accountID, amount int64, transactionType string) (*transaction.Transaction, error) {
	trans, err := r.q.InsertTransaction(ctx, postgresql.InsertTransactionParams{
		AccountID:       accountID,
		Amount:          amount,
		TransactionType: transactionType,
	})

	if err != nil {
		return nil, err
	}

	t := trans.MapToEntity()

	return &t, nil
}

func (r Repo) GetAccount(ctx context.Context, userID, accountID int64) (*transaction.Account, error) {
	account, err := r.q.GetAccount(ctx, postgresql.GetAccountParams{
		ID:     accountID,
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	a := account.MapToEntity()

	return &a, nil
}

func (r Repo) UpdateAccount(ctx context.Context, accountID, balance int64) error {
	_, err := r.q.UpdateBalanceAccount(ctx, postgresql.UpdateBalanceAccountParams{
		Balance:   balance,
		AccountID: accountID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) ListTransactions(ctx context.Context, userID int64) ([]transaction.Transaction, error) {
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

func (r Repo) ListTransactionsByAccount(ctx context.Context, userId, accountID int64) ([]transaction.Transaction, error) {
	ts, err := r.q.ListTransactionsByAccount(ctx, postgresql.ListTransactionsByAccountParams{
		ID:        userId,
		AccountID: accountID,
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
