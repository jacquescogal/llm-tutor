package handlers

import (
	"bff/internal/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
    controller *controllers.AuthenticationController
}

func NewAuthenticationHandler(controller *controllers.AuthenticationController) *AuthenticationHandler {
    return &AuthenticationHandler{
        controller: controller,
    }
}

func (h *AuthenticationHandler) CreateUser(c *gin.Context) {
    // Call the controller to handle CreateUser logic
    err := h.controller.CreateUser(c)
    fmt.Println("err", err)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *AuthenticationHandler) CreateSession(c *gin.Context) {
    // Call the controller to handle CreateSession logic
    sessionID, err := h.controller.CreateSession(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Set the session ID  as cookie
    c.SetCookie("session_id", sessionID, 3600, "/", "localhost", false, true)
    c.JSON(http.StatusOK, gin.H{"message": "Session created successfully"})
}

func (h *AuthenticationHandler) DeleteSession(c *gin.Context) {
    // Call the controller to handle DeleteSession logic
    fmt.Println("DeleteSession")
    err := h.controller.DeleteSession(c)
    fmt.Println("done")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}
