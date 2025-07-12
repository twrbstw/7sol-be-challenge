package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (a AuthHandler) RegisterRoutes(r fiber.Router) {
	r.Post("/register", a.Register)
	r.Post("/login", a.Login)
}

func (a AuthHandler) Register(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}

func (a AuthHandler) Login(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("")
}
