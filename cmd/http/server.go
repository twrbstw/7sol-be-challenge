package http

import (
	"context"
	"log"
	middleware "seven-solutions-challenge/internal/adapters/inbound"
	"seven-solutions-challenge/internal/adapters/inbound/http/routes"
	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/adapters/outbound/workers"
	"seven-solutions-challenge/internal/app/ports"
	"seven-solutions-challenge/internal/domain"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func StartHttp(ctx context.Context, cfg domain.Configs, client d.DatabaseConnection, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	loggerMiddleware := middleware.NewLoggerMiddleware(cfg.LoggerConfig)
	authMiddleware := middleware.NewAuthMiddleware(cfg.AppConfig)

	userRepo := initRepositories(client)

	app := fiber.New()
	app.Use(loggerMiddleware)
	app.Route("/auth", routes.RegisterAuthRoutes(userRepo, cfg.AppConfig))

	privateRoutes := app.Group("/api", authMiddleware) //add auth middleware
	privateRoutes.Route("/user", routes.RegisterPrivateRoutes(userRepo))

	workersList := initWorkers(userRepo)
	startWorkers(ctx, wg, workersList)

	startAppWithGracefulShutdown(app, ctx)
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

func startAppWithGracefulShutdown(app *fiber.App, ctx context.Context) {
	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Println("Fiber server stopped:", err)
		}
	}()

	log.Println("HTTP server running on :8080")

	<-ctx.Done() // wait for cancel from main
	log.Println("Shutting down HTTP server and worker(s)...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Fiber shutdown error: %v", err)
	}

	log.Println("HTTP server and worker(s) shutdown complete.")
}
