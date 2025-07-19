package routes

import (
	"seven-solutions-challenge/src/handlers"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterPrivateRoutes(userRepo repositories.IUserRepo) func(r fiber.Router) {
	userService := services.NewUserService(userRepo)

	return func(r fiber.Router) {
		userHandler := handlers.NewUserHandler(userService)
		userHandler.RegisterRoutes(r)
	}
}
