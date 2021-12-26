package postgresql

import (
	"minibank/account"
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
		ID:              t.ID,
		AccountID:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}

func (a Account) MapToEntity() account.Account {
	return account.Account{
		ID:        a.ID,
		UserID:    a.UserID,
		BankID:    a.BankID,
		Name:      a.Name,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}

}
