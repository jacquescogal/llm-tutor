package handlers


import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModuleHandler struct {
    controller *controllers.ModuleController
}

func NewModuleHandler(controller *controllers.ModuleController) *ModuleHandler {
    return &ModuleHandler{
        controller: controller,
    }
}

func (h *ModuleHandler) CreateModule(c *gin.Context) {
	// Call the controller to handle CreateModule logic
	err := h.controller.CreateModule(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Module created successfully"})
}
func (h *ModuleHandler) GetPublicModules(c *gin.Context) {
	res, err := h.controller.GetPublicModules(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ModuleHandler) GetPrivateModulesByUserId(c *gin.Context) {
	res, err := h.controller.GetPrivateModulesByUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ModuleHandler) GetFavouriteModulesByUserId(c *gin.Context) {
	res, err := h.controller.GetFavouriteModulesByUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ModuleHandler) GetModuleById(c *gin.Context) {
	res, err := h.controller.GetModuleById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ModuleHandler) GetModulesBySubjectId(c *gin.Context) {
	// Call the controller to handle GetModulesBySubjectId logic
	res, err := h.controller.GetModulesBySubjectId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ModuleHandler) GetModulesByNameSearch(c *gin.Context) {
	// Call the controller to handle GetModulesByNameSearch logic
	res, err := h.controller.GetModulesByNameSearch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *ModuleHandler) UpdateModule(c *gin.Context) {
	// Call the controller to handle UpdateModule logic
	err := h.controller.UpdateModule(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Module updated successfully"})
}

func (h *ModuleHandler) DeleteModule(c *gin.Context) {
	// Call the controller to handle DeleteModule logic
	err := h.controller.DeleteModule(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Module deleted successfully"})
}

func (h *ModuleHandler) SetUserModuleFavourite(c *gin.Context) {
	// Call the controller to handle SetUserModuleFavourite logic
	err := h.controller.SetUserModuleFavourite(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User module favourite set successfully"})
}
