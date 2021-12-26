package postgresql

import (
	"minibank/auth"
	"minibank/transaction"
)

func (u User) MapToEntity() auth.User {
	return auth.User{
		ID:        u.ID,
		Username:  u.Username,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (t Transaction) MapToEntity() transaction.Transaction {
	return transaction.Transaction{
		ID:            t.ID,
		FromAccountID: t.FromAccountID,
		ToAccountID:   t.ToAccountID,
		Amount:        t.Amount,
		Type:          t.Type,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
	}
}
