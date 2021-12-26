package usecase

import (
	"context"

	"minibank/transaction"
)

type TransUseCase struct {
	transRepo transaction.Repository
}

func (tu *TransUseCase) CreateDepositTransaction(ctx context.Context, userID, accountID, amount int64) (*transaction.Transaction, error) {
	//tim account
	_, err := tu.transRepo.GetAccount(ctx, userID, accountID)
	if err != nil {
		return nil, err
	}

	// tao transaction
	tran, err := tu.transRepo.CreateTransaction(ctx, userID, accountID, amount, "deposit")
	if err != nil {
		return nil, err
	}

	// update balance account
	if err = tu.transRepo.UpdateAccount(ctx, accountID, amount); err != nil {
		return nil, err
	}

	return tran, nil
}

func (tu *TransUseCase) CreateWithdrawTransaction(ctx context.Context, userID, accountID, amount int64) (*transaction.Transaction, error) {
	//TODO implement me
	panic("implement me")
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
