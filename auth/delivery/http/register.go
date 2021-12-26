package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/auth"
)

func RegisterHTTPEndpoints(router *fiber.App, uc auth.UseCase) {
	h := NewHandler(uc)

	endpoint := router.Group("/api/auth")
	{
		endpoint.Post("/signup", h.SignUp)
		endpoint.Post("/signin", h.SignIn)
	}
}
