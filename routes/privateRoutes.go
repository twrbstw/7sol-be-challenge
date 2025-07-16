package routes

import (
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/handlers"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterPrivateRoutes(db database.DatabaseConnection) func(r fiber.Router) {
	userRepo := repositories.NewUserRepo(db)
	userService := services.NewUserService(userRepo)

	return func(r fiber.Router) {
		userHandler := handlers.NewUserHandler(userService)
		userHandler.RegisterRoutes(r)
	}
}
