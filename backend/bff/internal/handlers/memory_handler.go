package handlers


import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemoryHandler struct {
    controller *controllers.MemoryController
}

func NewMemoryHandler(controller *controllers.MemoryController) *MemoryHandler {
    return &MemoryHandler{
        controller: controller,
    }
}

func (h *MemoryHandler) CreateMemory(c *gin.Context) {
	// Call the controller to handle CreateMemory logic
	err := h.controller.CreateMemory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory created successfully"})
}

func (h *MemoryHandler) GetMemoryById(c *gin.Context) {
	// Call the controller to handle GetMemoryById logic
	res, err := h.controller.GetMemoryById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *MemoryHandler) GetMemoriesByDocId(c *gin.Context) {
	// Call the controller to handle GetMemoriesByDocId logic
	res, err := h.controller.GetMemoriesByDocId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *MemoryHandler) GetMemoriesByMemoryTitleSearch(c *gin.Context) {
	// Call the controller to handle GetMemoriesByMemoryTitleSearch logic
	res, err := h.controller.GetMemoriesByMemoryTitleSearch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *MemoryHandler) UpdateMemory(c *gin.Context) {
	// Call the controller to handle UpdateMemory logic
	err := h.controller.UpdateMemory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory updated successfully"})
}

func (h *MemoryHandler) DeleteMemory(c *gin.Context) {
	// Call the controller to handle DeleteMemory logic
	err := h.controller.DeleteMemory(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory deleted successfully"})
}


