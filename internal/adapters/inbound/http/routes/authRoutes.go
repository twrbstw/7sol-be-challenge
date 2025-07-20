package routes

import (
	handler "seven-solutions-challenge/internal/adapters/inbound/http/handlers"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(userRepo repositories.IUserRepo, appCfg domain.AppConfig) func(r fiber.Router) {
	authService := services.NewAuthService(userRepo, appCfg)

	return func(r fiber.Router) {
		authHandler := handler.NewAuthHandler(authService)
		authHandler.RegisterRoutes(r)
	}
}
