package main

import (
	"bff/internal/controllers"
	"bff/internal/handlers"
	"bff/internal/proto/authenticator"
	// "bff/internal/proto/subject"
	"bff/internal/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
    r := gin.Default()
	// CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{os.Getenv("CORS_ALLOWED_ORIGINS")}, // Your frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	authServiceHost := os.Getenv("AUTH_SERVICE_HOST")
	authServicePort := os.Getenv("AUTH_SERVICE_PORT")
    conn, err := grpc.NewClient(fmt.Sprintf("%s:%s",authServiceHost,authServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect to gRPC server: %v", err)
    }
	// memoryServiceHost := os.Getenv("MEMORY_SERVICE_HOST")
	// memoryServicePort := os.Getenv("MEMORY_SERVICE_PORT")
	// memoryConn, err := grpc.NewClient(fmt.Sprintf("%s:%s",memoryServiceHost,memoryServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to gRPC server: %v", err)
	// }

    authService := services.NewAuthenticationService(authenticator.NewUserServiceClient(conn))
	// memoryService := services.NewMemoryService(memory.NewMemoryServiceClient(memoryConn))
	// documentService := services.NewDocumentService(document.NewDocServiceClient(memoryConn))
	// moduleService := services.NewModuleService(module.NewModuleServiceClient(memoryConn))
	// questionService := services.NewQuestionService(question.NewQuestionServiceClient(memoryConn))
	// subjectService := services.NewSubjectService(subject.NewSubjectServiceClient(memoryConn))
	// s3DB := db.NewS3UploadClient()

	authController := controllers.NewAuthenticationController(authService)
	// memoryController := controllers.NewMemoryController(memoryService)
	// documentController := controllers.NewDocumentController(s3DB,documentService)
	// moduleController := controllers.NewModuleController(moduleService)
	// questionController := controllers.NewQuestionController(questionService)
	// subjectController := controllers.NewSubjectController(subjectService)

    authHandler := handlers.NewAuthenticationHandler(authController)
	// memoryHandler := handlers.NewMemoryHandler(memoryController)
	// documentHandler := handlers.NewDocumentHandler(documentController)
	// moduleHandler := handlers.NewModuleHandler(moduleController)
	// questionHandler := handlers.NewQuestionHandler(questionController)
	// subjectHandler := handlers.NewSubjectHandler(subjectController)

	// Auth Routes
    r.POST("/user", authHandler.CreateUser)
    r.POST("/session", authHandler.CreateSession)
    r.DELETE("/session", handlers.SessionMiddleware(authService), authHandler.DeleteSession)

	// Subject Routes
	// r.POST("/subject", handlers.SessionMiddleware(authService), subjectHandler.CreateSubject)
	// r.GET("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectById)
	// r.GET("/user/subject", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectsByUserID)
	// r.POST("/search/subject", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectsByNameSearch)
	// r.PUT("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.UpdateSubject)
	// r.DELETE("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.DeleteSubject)

    // Run the server
	servicePort := os.Getenv("BFF_SERVICE_PORT")
    if err := r.Run(fmt.Sprintf(":%s",servicePort)); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}

