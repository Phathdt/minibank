package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/transaction"
)

func RegisterHTTPEndpoints(router *fiber.App, uc transaction.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/transactions")
	{
		authEndpoints.Get("", h.ListTransactions)
	}
}
