package handlers

import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DocumentHandler struct {
    controller *controllers.DocumentController
}

func NewDocumentHandler(controller *controllers.DocumentController) *DocumentHandler {
    return &DocumentHandler{
        controller: controller,
    }
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	// Call the controller to handle CreateDocument logic
	err := h.controller.CreateDocument(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Document created successfully"})
}

func (h *DocumentHandler) GetDocumentById(c *gin.Context) {
	// Call the controller to handle GetDocumentById logic
	res, err := h.controller.GetDocumentById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *DocumentHandler) GetDocumentsByModuleId(c *gin.Context) {
	// Call the controller to handle GetDocumentsByModuleId logic
	res, err := h.controller.GetDocumentsByModuleId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *DocumentHandler) GetDocumentsByNameSearch(c *gin.Context) {
	// Call the controller to handle GetDocumentsByNameSearch logic
	res, err := h.controller.GetDocumentsByNameSearch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *DocumentHandler) UpdateDocument(c *gin.Context) {
	// Call the controller to handle UpdateDocument logic
	err := h.controller.UpdateDocument(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Document updated successfully"})
}

func (h *DocumentHandler) DeleteDocument(c *gin.Context) {
	// Call the controller to handle DeleteDocument logic
	err := h.controller.DeleteDocument(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}
