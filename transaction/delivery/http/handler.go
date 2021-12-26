package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/transaction"
)

type Handler struct {
	useCase transaction.UseCase
}

func NewHandler(useCase transaction.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) ListTransactions(c *fiber.Ctx) error {
	transactions, err := h.useCase.ListTransactions(c.UserContext(), 1)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg":  "OK",
		"data": transactions,
	})
}
