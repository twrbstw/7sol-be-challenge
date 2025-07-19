package routes

import (
	"seven-solutions-challenge/src/handlers"
	"seven-solutions-challenge/src/models"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(userRepo repositories.IUserRepo, appCfg models.AppConfig) func(r fiber.Router) {
	authService := services.NewAuthService(userRepo, appCfg)

	return func(r fiber.Router) {
		authHandler := handlers.NewAuthHandler(authService)
		authHandler.RegisterRoutes(r)
	}
}
