package usecase

import (
	"context"

	"minibank/account"
)

type AccUseCase struct {
	accRepo account.Repository
}

func (au AccUseCase) UpdateAccount(ctx context.Context, userID, accountID int64, name string) (*account.Account, error) {
	acc, err := au.accRepo.GetAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if acc.UserID != userID {
		return nil, account.ErrNotYourAccount
	}

	return au.accRepo.UpdateAccount(ctx, accountID, name)
}

func (au AccUseCase) CreateAccount(ctx context.Context, userID, bankID int64, name string) (*account.Account, error) {
	return au.accRepo.CreateAccount(ctx, userID, bankID, name)
}

func (au AccUseCase) ListAccounts(ctx context.Context, userID int64) ([]account.Account, error) {
	return au.accRepo.ListAccounts(ctx, userID)
}

func NewAccUseCase(accRepo account.Repository) *AccUseCase {
	return &AccUseCase{accRepo: accRepo}
}
