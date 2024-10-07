package handlers

import (
	"bff/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware(authService *services.AuthenticationService) gin.HandlerFunc {
    return func(c *gin.Context) {
        sessionID, err := c.Cookie("session_id")
        if err!=nil{
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }


        userSession, err := authService.AuthenticateSession(sessionID)
        fmt.Println("SessionMiddleware done", userSession, err)
        if err != nil  {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        c.Set("user_session", userSession)
        c.Next()
    }
}
