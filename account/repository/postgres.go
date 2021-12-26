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

func NewAccountRepo(db *sql.DB) *Repo {
	return &Repo{q: postgresql.New(db)}
}
