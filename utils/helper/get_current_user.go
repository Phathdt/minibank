package helper

import (
	"github.com/gofiber/fiber/v2"
	"minibank/auth"
)

func GetCurrentUser(c *fiber.Ctx) *auth.User {
	currentUser := c.Locals("currentUser").(*auth.User)

	return currentUser
}
