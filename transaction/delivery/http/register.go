package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/transaction"
)

func RegisterHTTPEndpoints(router *fiber.App, uc transaction.UseCase) {
	h := NewHandler(uc)

	endpoint := router.Group("/transactions")
	{
		endpoint.Get("", h.ListTransactions)
		endpoint.Post("", h.CreateTransaction)
	}
}
