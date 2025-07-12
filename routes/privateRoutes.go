package routes

import (
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/handlers"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterPrivateRoutes(db database.DatabaseConnection) func(r fiber.Router) {
	userService := services.NewUserService(db)

	return func(r fiber.Router) {
		userHandler := handlers.NewUserHandler(userService)
		userHandler.RegisterRoutes(r)
	}
}
