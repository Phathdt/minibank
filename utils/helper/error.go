package helper

import (
	"github.com/gofiber/fiber/v2"
	"minibank/utils/validator"
)

func SimpleError(c *fiber.Ctx, err error) error {
	resp := validator.ToErrResponse(err)

	if resp == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(resp)
}
