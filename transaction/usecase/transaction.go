package usecase

import (
	"context"

	"minibank/account"
	"minibank/transaction"
)

type TransUseCase struct {
	transRepo transaction.Repository
	accRepo   account.Repository
}

func NewTransUseCase(transRepo transaction.Repository, accRepo account.Repository) *TransUseCase {
	return &TransUseCase{transRepo: transRepo, accRepo: accRepo}
}

func (tu *TransUseCase) CreateDepositTransaction(ctx context.Context, userID, accountID, amount int64) (*transaction.Transaction, error) {
	_, err := tu.accRepo.GetAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	// TODO: wrap transaction db
	tran, err := tu.transRepo.CreateTransaction(ctx, accountID, amount, "deposit")
	if err != nil {
		return nil, err
	}

	if err = tu.accRepo.UpdateAccount(ctx, accountID, amount); err != nil {
		return nil, err
	}

	return tran, nil
}

func (tu *TransUseCase) CreateWithdrawTransaction(ctx context.Context, userID, accountID, amount int64) (*transaction.Transaction, error) {
	acc, err := tu.accRepo.GetAccount(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if acc.UserID != userID {
		return nil, account.ErrNotYourAccount
	}

	if acc.Balance < amount {
		return nil, transaction.ErrAccountBalance
	}

	tran, err := tu.transRepo.CreateTransaction(ctx, accountID, amount, "withdraw")
	if err != nil {
		return nil, err
	}

	if err := tu.accRepo.UpdateAccount(ctx, accountID, -amount); err != nil {
		return nil, err
	}

	return tran, nil
}

func (tu *TransUseCase) ListTransactions(ctx context.Context, userID int64) ([]transaction.Transaction, error) {
	return tu.transRepo.ListTransactions(ctx, userID)
}

func (tu *TransUseCase) ListTransactionsByAccount(ctx context.Context, userID, accountID int64) ([]transaction.Transaction, error) {
	return tu.transRepo.ListTransactionsByAccount(ctx, userID, accountID)
}
