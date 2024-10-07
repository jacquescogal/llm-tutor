package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"authentication_service/internal/db"
	"authentication_service/internal/cache"
	"authentication_service/internal/controller"
	"authentication_service/internal/repository"
	"authentication_service/internal/handler"
	pb "authentication_service/internal/proto/authenticator"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// Initialize Redis connection
	redisClient, err := cache.NewRedis()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer redisClient.Close()

	// Initialize MySQL connection
	mysqlDB := db.NewDatabase()
	defer mysqlDB.Close()

	// Initialize repositories and controllers
	userRepo := repository.NewUserRepo(mysqlDB.Conn)
	distributedLock := cache.NewDistributedLock(redisClient.Client)
	sessionStore := cache.NewSessionStore(redisClient.Client)
	authController := controller.NewAuthenticationController(userRepo, distributedLock, sessionStore)

	// Initialize gRPC server and register services
	grpcServer := grpc.NewServer()
	userHandler := handler.NewAuthenticationHandler(authController)

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	// Start gRPC server and listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	log.Println("gRPC server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}