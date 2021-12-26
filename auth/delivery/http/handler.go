package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"minibank/auth"
	"minibank/utils/helper"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

type CreateUserDTO struct {
	Email    string `form:"email" validate:"required,min=6,max=32"`
	Password string `form:"password" validate:"required,min=6,max=32"`
}

func (h *Handler) SignUp(c *fiber.Ctx) error {
	data := new(CreateUserDTO)

	if err := c.BodyParser(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := validator.New().Struct(data); err != nil {
		return helper.SimpleError(c, err)
	}

	if err := h.useCase.SignUp(c.UserContext(), data.Email, data.Password); err != nil {
		return helper.SimpleError(c, err)
	}

	return c.Status(200).JSON(&fiber.Map{
		"msg": "OK",
	})
}
