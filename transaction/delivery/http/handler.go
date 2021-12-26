package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"minibank/transaction"
	"minibank/utils/helper"
)

type Handler struct {
	useCase transaction.UseCase
}

func NewHandler(useCase transaction.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) ListTransactions(c *fiber.Ctx) error {
	user := helper.GetCurrentUser(c)
	transactions, err := h.useCase.ListTransactions(c.UserContext(), user.ID)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg":  "OK",
		"data": transactions,
	})
}

type CreateTransactionDTO struct {
	AccountID       int64  `json:"account_id" form:"account_id" validate:"required"`
	Amount          int64  `json:"amount" form:"amount" validate:"required"`
	TransactionType string `json:"transaction_type" form:"transaction_type" validate:"oneof=deposit withdraw"`
}

func (h *Handler) CreateTransaction(c *fiber.Ctx) error {
	data := new(CreateTransactionDTO)

	if err := c.BodyParser(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := validator.New().Struct(data); err != nil {
		return helper.SimpleError(c, err)
	}

	user := helper.GetCurrentUser(c)

	var t *transaction.Transaction
	var err error

	if data.TransactionType == "deposit" {
		t, err = h.useCase.CreateDepositTransaction(c.UserContext(), user.ID, data.AccountID, data.Amount)
	} else {
		t, err = h.useCase.CreateWithdrawTransaction(c.UserContext(), user.ID, data.AccountID, data.Amount)
	}

	if err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg":  "OK",
		"data": t,
	})
}
