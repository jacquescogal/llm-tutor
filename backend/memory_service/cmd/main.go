package main

import (
	"log"
	"memory_core/internal/cache"
	"memory_core/internal/controller"
	"memory_core/internal/db"
	"memory_core/internal/handler"
	"memory_core/internal/producer"
	"memory_core/internal/proto/document"
	"memory_core/internal/proto/memory"
	"memory_core/internal/proto/module"
	"memory_core/internal/proto/question"
	"memory_core/internal/proto/subject"
	"memory_core/internal/proto/vector"
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
	
	// Initialize Kafka producer
	kafkaProducer := producer.NewKafkaProducer()
	// Initialize repositories and controllers
	cacheStore := cache.NewCacheStore(redisClient.Client)
	subjectRepo := repository.NewSubjectRepository()
	moduleRepo := repository.NewModuleRepository()
	userSubjectMapRepo := repository.NewUserSubjectMapRepository(cacheStore)
	userModuleMapRepo := repository.NewUserModuleMapRepository(cacheStore)
	subjectModuleMapRepo := repository.NewSubjectModuleMapRepository(cacheStore)
	docRepo := repository.NewDocRepository()
	// vector repo
	vectorRepo := repository.NewVectorRepository(db.NewWeaviate().Conn)

	subjectController := controller.NewSubjectController(mysqlDB.Conn, subjectRepo, userSubjectMapRepo,subjectModuleMapRepo)
	moduleController := controller.NewModuleController(mysqlDB.Conn, moduleRepo, userModuleMapRepo)
	docController := controller.NewDocController(mysqlDB.Conn, docRepo, moduleRepo, userModuleMapRepo, kafkaProducer)
	questionController := controller.NewQuestionController(mysqlDB.Conn, repository.NewQuestionRepository(), docRepo, moduleRepo, userModuleMapRepo)
	memoryController := controller.NewMemoryController(mysqlDB.Conn, repository.NewMemoryRepository(), userModuleMapRepo, moduleRepo, docRepo, vectorRepo)
	vectorController := controller.NewVectorController(mysqlDB.Conn, vectorRepo,subjectModuleMapRepo)


	// Initialize gRPC server and handlers
	grpcServer := grpc.NewServer()
	subjectHandler := handler.NewSubjectHandler(subjectController)
	moduleHandler := handler.NewModuleHandler(moduleController)
	docHandler := handler.NewDocHandler(docController)
	questionHandler := handler.NewQuestionHandler(questionController)
	memoryHandler := handler.NewMemoryHandler(memoryController)
	vectorHandler := handler.NewVectorHandler(vectorController)

	// Register services
	subject.RegisterSubjectServiceServer(grpcServer, subjectHandler)
	module.RegisterModuleServiceServer(grpcServer, moduleHandler)
	document.RegisterDocServiceServer(grpcServer, docHandler)
	question.RegisterQuestionServiceServer(grpcServer, questionHandler)
	memory.RegisterMemoryServiceServer(grpcServer, memoryHandler)
	vector.RegisterVectorServiceServer(grpcServer, vectorHandler)

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