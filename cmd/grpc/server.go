package grpc

import (
	"context"
	"log"
	"net"
	"sync"

	"seven-solutions-challenge/internal/adapters/inbound"
	"seven-solutions-challenge/internal/adapters/inbound/grpc"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/adapters/outbound/hasher"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/domain"
	"seven-solutions-challenge/proto/userpb"

	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"

	g "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpc(ctx context.Context, cfg domain.Configs, client d.DatabaseConnection, wg *sync.WaitGroup) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := g.NewServer(
		g.UnaryInterceptor(
			inbound.NewAuthInterceptor(cfg.AppConfig),
		),
	)

	userRepo := repositories.NewUserRepo(client)
	bcryptHasher := hasher.NewBcryptHasher()
	userService := services.NewUserService(userRepo, bcryptHasher)
	userpb.RegisterUserServiceServer(grpcServer, &grpc.UserServiceServer{
		UserService: userService,
	})

	reflection.Register(grpcServer)

	serverErr := make(chan error, 1)
	go func() {
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			serverErr <- err
		}
	}()

	select {
	case <-ctx.Done():
		// context cancelled -> shutdown
		log.Println("Shutting down gRPC server...")
		grpcServer.GracefulStop()
		log.Println("gRPC server shutdown complete.")

	case err := <-serverErr:
		// server crashed
		log.Printf("gRPC server stopped with error: %v", err)
	}
}
