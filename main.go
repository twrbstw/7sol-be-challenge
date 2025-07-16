package main

import (
	"context"
	"seven-solutions-challenge/routes"
	"seven-solutions-challenge/src/config"
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {

	ctx := context.Background()
	cfg := config.LoadDefaultConfig()

	client := database.NewDatabaseClient(ctx, cfg.DbConfig)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	loggerMiddleware := middleware.NewLoggerMiddleware(cfg.LoggerConfig)

	app := fiber.New()
	app.Use(loggerMiddleware)
	app.Route("/auth", routes.RegisterAuthRoutes(client))

	privateRoutes := app.Group("/api") //add auth middleware
	privateRoutes.Route("/user", routes.RegisterPrivateRoutes(client))

	app.Listen(":3000")
}
