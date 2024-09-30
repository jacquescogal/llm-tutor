package main

import (
	"bff/internal/controllers"
	"bff/internal/handlers"
	"bff/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if present
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
    r := gin.Default()
    authService := services.NewAuthenticationService()
	authController := controllers.NewAuthenticationController(authService)
    authHandler := handlers.NewAuthenticationHandler(authController)

    r.POST("/api/auth/create_user", authHandler.CreateUser)
    r.POST("/api/auth/create_session", authHandler.CreateSession)
    r.DELETE("/api/auth/delete_session", handlers.SessionMiddleware(authService), authHandler.DeleteSession)

    // Run the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Could not run server: %v", err)
    }
}
