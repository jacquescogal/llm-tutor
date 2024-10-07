package handlers


import (
	"bff/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
    controller *controllers.QuestionController
}

func NewQuestionHandler(controller *controllers.QuestionController) *QuestionHandler {
    return &QuestionHandler{
        controller: controller,
    }
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	// Call the controller to handle CreateQuestion logic
	err := h.controller.CreateQuestion(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question created successfully"})
}

func (h *QuestionHandler) GetQuestionById(c *gin.Context) {
	// Call the controller to handle GetQuestionById logic
	res, err := h.controller.GetQuestionById(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) GetQuestionsByDocId(c *gin.Context) {
	// Call the controller to handle GetQuestionsByDocId logic
	res, err := h.controller.GetQuestionsByDocId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) GetQuestionsByQuestionTitleSearch(c *gin.Context) {
	// Call the controller to handle GetQuestionsByQuestionTitleSearch logic
	res, err := h.controller.GetQuestionsByQuestionTitleSearch(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *QuestionHandler) UpdateQuestion(c *gin.Context) {
	// Call the controller to handle UpdateQuestion logic
	err := h.controller.UpdateQuestion(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question updated successfully"})
}

func (h *QuestionHandler) DeleteQuestion(c *gin.Context) {
	// Call the controller to handle DeleteQuestion logic
	err := h.controller.DeleteQuestion(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}
