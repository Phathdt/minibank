package http

import (
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
