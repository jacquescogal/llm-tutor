package main

import (
	"log"
	"memory_core/internal/cache"
	"memory_core/internal/controller"
	"memory_core/internal/db"
	"memory_core/internal/handler"
	"memory_core/internal/proto/subject"
	"memory_core/internal/repository"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
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
	cacheStore := cache.NewCacheStore(redisClient.Client)
	subjectRepo := repository.NewSubjectRepository()
	userSubjectMapRepo := repository.NewUserSubjectMapRepository(cacheStore)
	subjectController := controller.NewSubjectController(mysqlDB.Conn, subjectRepo, userSubjectMapRepo)

	// Initialize gRPC server and register services
	grpcServer := grpc.NewServer()
	subjectHandler := handler.NewSubjectHandler(subjectController)

	subject.RegisterSubjectServiceServer(grpcServer, subjectHandler)

	// Start gRPC server and listen on port 50051
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}
	log.Println("gRPC server running on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC: %v", err)
	}
}