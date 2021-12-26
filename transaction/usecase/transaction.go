package usecase

import (
	"context"

	"minibank/transaction"
)

type TransUseCase struct {
	transRepo transaction.Repository
}

func NewTransUseCase(transRepo transaction.Repository) *TransUseCase {
	return &TransUseCase{transRepo: transRepo}
}

func (tu *TransUseCase) ListTransactions(ctx context.Context, userID int64) ([]transaction.Transaction, error) {
	return tu.transRepo.ListTransactions(ctx, userID)
}

func (tu *TransUseCase) ListTransactionsByAccount(ctx context.Context, userID, accountID int64) ([]transaction.Transaction, error) {
	return tu.transRepo.ListTransactionsByAccount(ctx, userID, accountID)
}
