package main

import (
	"context"
	"log"
	"net"

	"seven-solutions-challenge/internal/adapters/inbound/grpc"
	d "seven-solutions-challenge/internal/adapters/outbound/db/mongo"
	"seven-solutions-challenge/internal/adapters/outbound/db/mongo/repositories"
	"seven-solutions-challenge/internal/adapters/outbound/hasher"
	"seven-solutions-challenge/internal/app/services"
	"seven-solutions-challenge/internal/config"
	"seven-solutions-challenge/proto/userpb"

	g "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadDefaultConfig()
	client := d.NewDatabaseClient(ctx, cfg.DbConfig)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := g.NewServer()

	userRepo := repositories.NewUserRepo(client)
	bcryptHasher := hasher.NewBcryptHasher()
	userService := services.NewUserService(userRepo, bcryptHasher)
	userpb.RegisterUserServiceServer(grpcServer, &grpc.UserServiceServer{
		UserService: userService,
	})

	reflection.Register(grpcServer)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
