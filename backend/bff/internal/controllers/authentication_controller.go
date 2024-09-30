package controllers

import (
	auth "bff/internal/proto/authenticator"
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
    var request auth.CreateUserRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        return err
    }
    return c.authService.CreateUser(request.Username, request.Password)
}

func (c *AuthenticationController) CreateSession(ctx *gin.Context) (string, error) {
    var request auth.CreateSessionRequest
    if err := ctx.ShouldBindJSON(&request); err != nil {
        return "", err
    }
    return c.authService.CreateSession(request.Username, request.Password)
}

func (c *AuthenticationController) DeleteSession(ctx *gin.Context) error {
    sessionID, err := ctx.Cookie("session_id")
    fmt.Println(sessionID)
    if err != nil {
        return err
    }
    return c.authService.DeleteSession(sessionID)
}
