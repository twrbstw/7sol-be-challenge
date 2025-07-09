package main

import (
	"context"
	"seven-solutions-assignment/config"
	"seven-solutions-assignment/database"
)

// "github.com/gofiber/fiber/v2"

func main() {

	ctx := context.Background()
	cfg := config.LoadDefaultConfig()

	dbClient := database.NewDatabaseClient(ctx, cfg.DbConfig)
	// app := fiber.New()

	// app.Listen(":3000")
}
