package repository

import (
	"context"
	"database/sql"

	"minibank/account"
	"minibank/db/postgresql"
)

type Repo struct {
	q *postgresql.Queries
}

func (r Repo) ListAccounts(ctx context.Context, userID int64) ([]account.Account, error) {
	accs, err := r.q.ListAccounts(ctx, userID)
	if err != nil {
		return nil, err
	}

	accounts := make([]account.Account, len(accs))

	for i, acc := range accs {
		accounts[i] = acc.MapToEntity()
	}

	return accounts, nil
}

func (r Repo) GetAccount(ctx context.Context, accountID int64) (*account.Account, error) {
	acc, err := r.q.GetAccount(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, account.ErrAccountNotFound
		}

		return nil, err
	}

	a := acc.MapToEntity()

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
func NewAccountRepo(db *sql.DB) *Repo {
	return &Repo{q: postgresql.New(db)}
}
