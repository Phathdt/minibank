package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/account"
	"minibank/utils/helper"
)

type Handler struct {
	useCase account.UseCase
}

func NewHandler(useCase account.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) ListAccounts(c *fiber.Ctx) error {
	user := helper.GetCurrentUser(c)

	accounts, err := h.useCase.ListAccounts(c.UserContext(), user.ID)
	if err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg":  "OK",
		"data": accounts,
	})
}
