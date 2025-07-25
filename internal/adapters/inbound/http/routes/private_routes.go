package routes

import (
	handler "seven-solutions-challenge/internal/adapters/inbound/http/handlers"
	"seven-solutions-challenge/internal/adapters/outbound/hasher"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/app/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterPrivateRoutes(userRepo ports.IUserRepo) func(r fiber.Router) {
	bcryptHasher := hasher.NewBcryptHasher()
	userService := services.NewUserService(userRepo, bcryptHasher)

	return func(r fiber.Router) {
		userHandler := handler.NewUserHandler(userService)
		userHandler.RegisterRoutes(r)
	}
}
