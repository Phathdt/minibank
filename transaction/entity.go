package transaction

import "time"

type Transaction struct {
	ID              int64     `json:"id"`
	AccountID       int64     `json:"account_id"`
	Amount          int64     `json:"amount"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	BankID    int64     `json:"bank_id"`
	Name      string    `json:"name"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
