package http

import (
	"github.com/go-playground/validator/v10"
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

type CreateAccountDTO struct {
	BankID int64  `json:"bank_id" form:"bank_id" validate:"required"`
	Name   string `json:"name" form:"name" validate:"required"`
}

func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	data := new(CreateAccountDTO)

	if err := c.BodyParser(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := validator.New().Struct(data); err != nil {
		return helper.SimpleError(c, err)
	}

	user := helper.GetCurrentUser(c)

	acc, err := h.useCase.CreateAccount(c.UserContext(), user.ID, data.BankID, data.Name)
	if err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg":  "OK",
		"data": acc,
	})
}
