package main

import (
	"bff/internal/controllers"
	"bff/internal/db"
	"bff/internal/handlers"
	"bff/internal/proto/authenticator"
	"bff/internal/proto/document"
	"bff/internal/proto/generation"
	"bff/internal/proto/memory"
	"bff/internal/proto/module"
	"bff/internal/proto/question"
	"bff/internal/proto/subject"

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
	memoryServiceHost := os.Getenv("MEMORY_SERVICE_HOST")
	memoryServicePort := os.Getenv("MEMORY_SERVICE_PORT")
	memoryConn, err := grpc.NewClient(fmt.Sprintf("%s:%s",memoryServiceHost,memoryServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	generationServiceHost := os.Getenv("GEN_SERVICE_HOST")
	generationServicePort := os.Getenv("GEN_SERVICE_PORT")
	generationConn, err := grpc.NewClient(fmt.Sprintf("%s:%s",generationServiceHost,generationServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))

    authService := services.NewAuthenticationService(authenticator.NewUserServiceClient(conn))
	memoryService := services.NewMemoryService(memory.NewMemoryServiceClient(memoryConn))
	documentService := services.NewDocumentService(document.NewDocServiceClient(memoryConn))
	moduleService := services.NewModuleService(module.NewModuleServiceClient(memoryConn))
	questionService := services.NewQuestionService(question.NewQuestionServiceClient(memoryConn))
	subjectService := services.NewSubjectService(subject.NewSubjectServiceClient(memoryConn))
	generationService := services.NewGenerationService(generation.NewGenerationServiceClient(generationConn))
	s3DB := db.NewS3UploadClient()

	authController := controllers.NewAuthenticationController(authService)
	memoryController := controllers.NewMemoryController(memoryService)
	documentController := controllers.NewDocumentController(s3DB,documentService)
	moduleController := controllers.NewModuleController(moduleService)
	questionController := controllers.NewQuestionController(questionService)
	subjectController := controllers.NewSubjectController(subjectService)
	generationController := controllers.NewGenerationController(generationService)

    authHandler := handlers.NewAuthenticationHandler(authController)
	memoryHandler := handlers.NewMemoryHandler(memoryController)
	documentHandler := handlers.NewDocumentHandler(documentController)
	moduleHandler := handlers.NewModuleHandler(moduleController)
	questionHandler := handlers.NewQuestionHandler(questionController)
	subjectHandler := handlers.NewSubjectHandler(subjectController)
	generationHandler := handlers.NewGenerationHandler(generationController)

	// Auth Routes
    r.POST("/user", authHandler.CreateUser)
    r.POST("/session", authHandler.CreateSession)
    r.DELETE("/session", handlers.SessionMiddleware(authService), authHandler.DeleteSession)

	// Generation Routes
	r.POST("/generate", handlers.SessionMiddleware(authService), generationHandler.CreateGeneration)

	// Subject Routes
	r.POST("/subject", handlers.SessionMiddleware(authService), subjectHandler.CreateSubject)
	r.GET("/public/subject", handlers.SessionMiddleware(authService), subjectHandler.GetPublicSubjects)
	r.GET("/private/subject", handlers.SessionMiddleware(authService), subjectHandler.GetPrivateSubjectsByUserId)
	r.GET("/favourite/subject", handlers.SessionMiddleware(authService), subjectHandler.GetFavouriteSubjectsByUserId)
	r.GET("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectById)
	r.GET("/user/subject", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectsByUserID)
	r.POST("/search/subject", handlers.SessionMiddleware(authService), subjectHandler.GetSubjectsByNameSearch)
	r.POST("/favourite/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.SetUserSubjectFavourite)
	r.PUT("/subject/:subject_id/module", handlers.SessionMiddleware(authService), subjectHandler.SetSubjectModuleMapping)
	r.PUT("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.UpdateSubject)
	r.DELETE("/subject/:subject_id", handlers.SessionMiddleware(authService), subjectHandler.DeleteSubject)

	// Module Routes
	r.POST("/module", handlers.SessionMiddleware(authService), moduleHandler.CreateModule)
	r.GET("/public/module", handlers.SessionMiddleware(authService), moduleHandler.GetPublicModules)
	r.GET("/private/module", handlers.SessionMiddleware(authService), moduleHandler.GetPrivateModulesByUserId)
	r.GET("/favourite/module", handlers.SessionMiddleware(authService), moduleHandler.GetFavouriteModulesByUserId)
	r.GET("/module/:module_id", handlers.SessionMiddleware(authService), moduleHandler.GetModuleById)
	r.GET("/subject/:subject_id/module", handlers.SessionMiddleware(authService), moduleHandler.GetModulesBySubjectId)
	r.POST("/search/module", handlers.SessionMiddleware(authService), moduleHandler.GetModulesByNameSearch)
	r.POST("/favourite/module/:module_id", handlers.SessionMiddleware(authService), moduleHandler.SetUserModuleFavourite)
	// r.POST("/search/subject/:subject_id/module", handlers.SessionMiddleware(authService), moduleHandler.GetSubjectModulesByNameSearch)
	r.PUT("/module/:module_id", handlers.SessionMiddleware(authService), moduleHandler.UpdateModule)
	r.DELETE("/module/:module_id", handlers.SessionMiddleware(authService), moduleHandler.DeleteModule)

	// Document Routes
	r.POST("/module/:module_id/document", handlers.SessionMiddleware(authService), documentHandler.CreateDocument)
	r.GET("/module/:module_id/document/:document_id", handlers.SessionMiddleware(authService), documentHandler.GetDocumentById)
	r.GET("/module/:module_id/document", handlers.SessionMiddleware(authService), documentHandler.GetDocumentsByModuleId)
	r.POST("/search/module/:module_id/document", handlers.SessionMiddleware(authService), documentHandler.GetDocumentsByNameSearch)
	r.PUT("/module/:module_id/document/:document_id", handlers.SessionMiddleware(authService), documentHandler.UpdateDocument)
	r.DELETE("/module/:module_id/document/:document_id", handlers.SessionMiddleware(authService), documentHandler.DeleteDocument)

	// Memory Routes
	r.POST("module/:module_id/document/:document_id/memory", handlers.SessionMiddleware(authService), memoryHandler.CreateMemory)
	r.GET("module/:module_id/document/:document_id/memory/:memory_id", handlers.SessionMiddleware(authService), memoryHandler.GetMemoryById)
	r.GET("module/:module_id/document/:document_id/memory", handlers.SessionMiddleware(authService), memoryHandler.GetMemoriesByDocId)
	r.POST("search/module/:module_id/document/:document_id/memory", handlers.SessionMiddleware(authService), memoryHandler.GetMemoriesByMemoryTitleSearch)
	r.PUT("module/:module_id/document/:document_id/memory/:memory_id", handlers.SessionMiddleware(authService), memoryHandler.UpdateMemory)
	r.DELETE("module/:module_id/document/:document_id/memory/:memory_id", handlers.SessionMiddleware(authService), memoryHandler.DeleteMemory)


	// Question Routes
	r.POST("module/:module_id/document/:document_id/question", handlers.SessionMiddleware(authService), questionHandler.CreateQuestion)
	r.GET("module/:module_id/document/:document_id/question/:question_id", handlers.SessionMiddleware(authService), questionHandler.GetQuestionById)
	r.GET("module/:module_id/document/:document_id/question", handlers.SessionMiddleware(authService), questionHandler.GetQuestionsByDocId)
	r.POST("search/module/:module_id/document/:document_id/question", handlers.SessionMiddleware(authService), questionHandler.GetQuestionsByQuestionTitleSearch)
	r.PUT("module/:module_id/document/:document_id/question/:question_id", handlers.SessionMiddleware(authService), questionHandler.UpdateQuestion)
	r.DELETE("module/:module_id/document/:document_id/question/:question_id", handlers.SessionMiddleware(authService), questionHandler.DeleteQuestion)

    // Run the server
	servicePort := os.Getenv("BFF_SERVICE_PORT")
    if err := r.Run(fmt.Sprintf(":%s",servicePort)); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}

