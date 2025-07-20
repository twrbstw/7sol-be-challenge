package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	middleware "seven-solutions-challenge/internal/adapters/inbound"
	"seven-solutions-challenge/internal/adapters/inbound/http/routes"
	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/adapters/outbound/workers"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/config"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	// ctx := context.Background()
	cfg := config.LoadDefaultConfig()

	client := d.NewDatabaseClient(ctx, cfg.DbConfig)
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
	startWorkers(ctx, &wg, workersList)

	startAppWithGracefulShutdown(app, cancel, &wg)
}

func initRepositories(client d.DatabaseConnection) ports.IUserRepo {
	userRepo := repositories.NewUserRepo(client)
	return userRepo
}

func initWorkers(userRepo ports.IUserRepo) []ports.IWorkers {
	var workersList []ports.IWorkers

	listUserWorker := workers.NewListUsersWorker(userRepo)

	workersList = append(workersList, listUserWorker)
	return workersList
}

func startWorkers(ctx context.Context, wg *sync.WaitGroup, workersList []ports.IWorkers) {
	for _, worker := range workersList {
		log.Println("Starting worker:", worker.GetWorkerName())
		go worker.Start(ctx, wg)
	}
}

func startAppWithGracefulShutdown(app *fiber.App, cancel context.CancelFunc, wg *sync.WaitGroup) {
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Println("Fiber server stopped:", err)
		}
	}()

	log.Println("Server is running on http://localhost:3000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	log.Println("Gracefully shutting down...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Fiber shutdown error: %v", err)
	}

	cancel()
	wg.Wait()
	log.Println("All workers stopped. Server shutdown complete.")
}
