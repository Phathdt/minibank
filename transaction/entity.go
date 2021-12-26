package transaction

import "time"

type Transaction struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	Type          string    `json:"type"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
