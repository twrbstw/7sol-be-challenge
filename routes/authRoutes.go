package routes

import (
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(db database.DatabaseConnection) func(r fiber.Router) {
	return func(r fiber.Router) {
		authHandler := handlers.NewAuthHandler()
		authHandler.RegisterRoutes(r)
	}
}
