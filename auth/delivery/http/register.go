package http

import (
	"github.com/gofiber/fiber/v2"
	"minibank/auth"
)

func RegisterHTTPEndpoints(router *fiber.App, uc auth.UseCase) {
	h := NewHandler(uc)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.Post("/signup", h.SignUp)
		authEndpoints.Post("/signin", h.SignIn)
	}
}
