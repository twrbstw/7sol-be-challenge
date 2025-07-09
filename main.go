package main

import (
	"context"
	"seven-solutions-challenge/routes"
	"seven-solutions-challenge/src/config"
	"seven-solutions-challenge/src/database"

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

	app := fiber.New()
	authrouter := app.Group("/api") // add middleware here
	authrouter.Route("/user", routes.RegisterRoutes(client))

	app.Listen(":3000")
}
