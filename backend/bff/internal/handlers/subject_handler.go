package handlers


import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
    controller *controllers.SubjectController
}

func NewSubjectHandler(controller *controllers.SubjectController) *SubjectHandler {
    return &SubjectHandler{
        controller: controller,
    }
}

func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	err := h.controller.CreateSubject(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject created successfully"})
}

func (h *SubjectHandler) GetPublicSubjects(c *gin.Context) {
	res, err := h.controller.GetPublicSubjects(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) GetPrivateSubjectsByUserId(c *gin.Context) {
	res, err := h.controller.GetPrivateSubjectsByUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) GetFavouriteSubjectsByUserId(c *gin.Context) {
	res, err := h.controller.GetFavouriteSubjectsByUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) GetSubjectById(c *gin.Context) {
	res, err := h.controller.GetSubjectByID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) GetSubjectsByUserID(c *gin.Context) {
	// Call the controller to handle GetSubjectsByUserID logic
	res, err := h.controller.GetSubjectsByUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) GetSubjectsByNameSearch(c *gin.Context) {
	// Call the controller to handle GetSubjectsByNameSearch logic
	res, err := h.controller.GetSubjectsByNameSearch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	// Call the controller to handle UpdateSubject logic
	err := h.controller.UpdateSubject(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject updated successfully"})
}

func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	// Call the controller to handle DeleteSubject logic
	err := h.controller.DeleteSubject(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

func (h *SubjectHandler) SetUserSubjectFavourite(c *gin.Context) {
	// Call the controller to handle SetUserSubjectFavourite logic
	err := h.controller.SetUserSubjectFavourite(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject favourited successfully"})
}