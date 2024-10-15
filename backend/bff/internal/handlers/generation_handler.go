package handlers

import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerationHandler struct {
    controller *controllers.GenerationController
}

func NewGenerationHandler(controller *controllers.GenerationController) *GenerationHandler {
    return &GenerationHandler{
        controller: controller,
    }
}

func (h *GenerationHandler) CreateGeneration(c *gin.Context) {
	// Call the controller to handle CreateGeneration logic
	res, err := h.controller.CreateGeneration(c)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}