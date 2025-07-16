package routes

import (
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/handlers"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(db database.DatabaseConnection) func(r fiber.Router) {
	userRepo := repositories.NewUserRepo(db)
	authService := services.NewAuthService(userRepo)

	return func(r fiber.Router) {
		authHandler := handlers.NewAuthHandler(authService)
		authHandler.RegisterRoutes(r)
	}
}
