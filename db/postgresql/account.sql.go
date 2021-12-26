// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package postgresql

import (
	"context"
)

const getAccount = `-- name: GetAccount :one
SELECT id, user_id, bank_id, name, balance, created_at, updated_at
FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BankID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_id, bank_id, name, balance, created_at, updated_at
FROM accounts
WHERE user_id = $1
ORDER BY id
`

func (q *Queries) ListAccounts(ctx context.Context, userID int64) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.BankID,
			&i.Name,
			&i.Balance,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBalanceAccount = `-- name: UpdateBalanceAccount :one
UPDATE
	accounts
SET
	balance = balance + $1
WHERE
	id = $2
RETURNING id, user_id, bank_id, name, balance, created_at, updated_at
`

type UpdateBalanceAccountParams struct {
	Balance   int64 `json:"balance"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) UpdateBalanceAccount(ctx context.Context, arg UpdateBalanceAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateBalanceAccount, arg.Balance, arg.AccountID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.BankID,
		&i.Name,
		&i.Balance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
