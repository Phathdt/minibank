package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/account"
)

func RegisterHTTPEndpoints(router *fiber.App, au account.UseCase) {
	h := NewHandler(au)

	authEndpoints := router.Group("/accounts")
	{
		authEndpoints.Get("", h.ListAccounts)
	}
}
