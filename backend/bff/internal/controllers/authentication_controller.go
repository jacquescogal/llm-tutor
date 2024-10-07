package controllers

import (
	"bff/internal/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
    authService *services.AuthenticationService
}

func NewAuthenticationController(authenticationService *services.AuthenticationService) *AuthenticationController {
    return &AuthenticationController{
        authService: authenticationService,
    }
}

func (c *AuthenticationController) CreateUser(ctx *gin.Context) error {
    username, password, ok := ctx.Request.BasicAuth()
    if !ok {
        return fmt.Errorf("basic auth header not found")
    }
    return c.authService.CreateUser(username, password)
}

func (c *AuthenticationController) CreateSession(ctx *gin.Context) (string, error) {
    // check the basic auth header for username and password
    username, password, ok := ctx.Request.BasicAuth()
    if !ok {
        return "", fmt.Errorf("basic auth header not found")
    }
    return c.authService.CreateSession(username, password)
}

func (c *AuthenticationController) DeleteSession(ctx *gin.Context) error {
    sessionID, err := ctx.Cookie("session_id")
    fmt.Println(sessionID)
    if err != nil {
        return err
    }
    return c.authService.DeleteSession(sessionID)
}
