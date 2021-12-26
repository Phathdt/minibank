package usecase

import (
	"context"

	"minibank/account"
)

type AccUseCase struct {
	accRepo account.Repository
}

func (au AccUseCase) ListAccounts(ctx context.Context, userID int64) ([]account.Account, error) {
	return au.accRepo.ListAccounts(ctx, userID)
}

func NewAccUseCase(accRepo account.Repository) *AccUseCase {
	return &AccUseCase{accRepo: accRepo}
}
