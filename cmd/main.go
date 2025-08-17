package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"seven-solutions-challenge/cmd/grpc"
	"seven-solutions-challenge/cmd/http"
	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"
	"seven-solutions-challenge/internal/config"
	"sync"
	"syscall"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.LoadDefaultConfig()

	client := d.NewDatabaseClient(ctx, cfg.DbConfig)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2) // HTTP + gRPC

	go grpc.StartGrpc(ctx, cfg, client, &wg)
	go http.StartHttp(ctx, cfg, client, cancel, &wg)

	// Listen for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Print("Shutdown signal received")
	cancel() // notify both servers

	wg.Wait() // wait for both to finish
	log.Println("All servers shut down gracefully")
}
