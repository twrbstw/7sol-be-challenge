package main

import (
	"context"
	"log"
	"seven-solutions-challenge/routes"
	"seven-solutions-challenge/src/config"
	"seven-solutions-challenge/src/database"
	"seven-solutions-challenge/src/middleware"
	repositories "seven-solutions-challenge/src/repositories/userRep"
	"seven-solutions-challenge/src/workers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// ctx, cancel := context.WithCancel(context.Background())
	ctx := context.Background()
	cfg := config.LoadDefaultConfig()

	client := database.NewDatabaseClient(ctx, cfg.DbConfig)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	loggerMiddleware := middleware.NewLoggerMiddleware(cfg.LoggerConfig)
	authMiddleware := middleware.NewAuthMiddleware(cfg.AppConfig)

	userRepo := initRepositories(client)

	app := fiber.New()
	app.Use(loggerMiddleware)
	app.Route("/auth", routes.RegisterAuthRoutes(userRepo, cfg.AppConfig))

	privateRoutes := app.Group("/api", authMiddleware) //add auth middleware
	privateRoutes.Route("/user", routes.RegisterPrivateRoutes(userRepo))

	workersList := initWorkers(userRepo)
	startWorkers(ctx, workersList)

	// cancel()
	// time.Sleep(1 * time.Second)
	app.Listen(":3000")
}

func initRepositories(client database.DatabaseConnection) repositories.IUserRepo {
	userRepo := repositories.NewUserRepo(client)
	return userRepo
}

func initWorkers(userRepo repositories.IUserRepo) []workers.IWorkers {
	var workersList []workers.IWorkers

	listUserWorker := workers.NewListUsersWorker(userRepo)

	workersList = append(workersList, listUserWorker)
	return workersList
}

func startWorkers(ctx context.Context, workersList []workers.IWorkers) {
	for _, worker := range workersList {
		log.Println("starting worker:", worker.GetWorkerName())
		go worker.Start(ctx)
	}
}
