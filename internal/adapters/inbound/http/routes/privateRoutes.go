package routes

import (
	handler "seven-solutions-challenge/internal/adapters/inbound/http/handlers"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/app/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterPrivateRoutes(userRepo repositories.IUserRepo) func(r fiber.Router) {
	userService := services.NewUserService(userRepo)

	return func(r fiber.Router) {
		userHandler := handler.NewUserHandler(userService)
		userHandler.RegisterRoutes(r)
	}
}
