package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/account"
)

func RegisterHTTPEndpoints(router fiber.Router, au account.UseCase) {
	h := NewHandler(au)

	authEndpoints := router.Group("/accounts")
	{
		authEndpoints.Get("", h.ListAccounts)
		authEndpoints.Post("", h.CreateAccount)
		authEndpoints.Put("/:id", h.UpdateAccount)
	}
}
