package routes

import (
	handler "seven-solutions-challenge/internal/adapters/inbound/http/handlers"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(userRepo ports.IUserRepo, appCfg domain.AppConfig) func(r fiber.Router) {
	authService := services.NewAuthService(userRepo, appCfg)

	return func(r fiber.Router) {
		authHandler := handler.NewAuthHandler(authService)
		authHandler.RegisterRoutes(r)
	}
}
