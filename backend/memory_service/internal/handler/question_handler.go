package handler

import (
	"context"
	"memory_core/internal/controller"
	"memory_core/internal/proto/question"
)

// QuestionHandler represents the handler for managing questions
type QuestionHandler struct {
	// Question Controller
	questionController *controller.QuestionController
	question.UnimplementedQuestionServiceServer
}

// NewQuestionHandler creates a new QuestionHandler
func NewQuestionHandler(questionController *controller.QuestionController) *QuestionHandler {
	return &QuestionHandler{questionController: questionController}
}

// CreateQuestion creates a new question
func (questionHandler *QuestionHandler) CreateQuestion(ctx context.Context, createQuestionRequest *question.CreateQuestionRequest) (*question.CreateQuestionResponse, error) {
	return questionHandler.questionController.CreateQuestion(ctx, createQuestionRequest)
}

// GetQuestionById retrieves a question by question_id
func (questionHandler *QuestionHandler) GetQuestionById(ctx context.Context, getQuestionByIdRequest *question.GetQuestionByIdRequest) (*question.GetQuestionByIdResponse, error) {
	return questionHandler.questionController.GetQuestionById(ctx, getQuestionByIdRequest)
}

// GetQuestionsByDocId retrieves all questions associated with a specific doc_id
func (questionHandler *QuestionHandler) GetQuestionsByDocId(ctx context.Context, getQuestionsByDocIdRequest *question.GetQuestionsByDocIdRequest) (*question.GetQuestionsByDocIdResponse, error) {
	return questionHandler.questionController.GetQuestionsByDocId(ctx, getQuestionsByDocIdRequest)
}

// GetQuestionsByQuestionTitleSearch retrieves questions by a title search
func (questionHandler *QuestionHandler) GetQuestionsByQuestionTitleSearch(ctx context.Context, getQuestionsByQuestionTitleSearchRequest *question.GetQuestionsByQuestionTitleSearchRequest) (*question.GetQuestionsByQuestionTitleSearchResponse, error) {
	return questionHandler.questionController.GetQuestionsByQuestionTitleSearch(ctx, getQuestionsByQuestionTitleSearchRequest)
}

// UpdateQuestion updates a question
func (questionHandler *QuestionHandler) UpdateQuestion(ctx context.Context, updateQuestionRequest *question.UpdateQuestionRequest) (*question.UpdateQuestionResponse, error) {
	return questionHandler.questionController.UpdateQuestion(ctx, updateQuestionRequest)
}

// DeleteQuestion deletes a question
func (questionHandler *QuestionHandler) DeleteQuestion(ctx context.Context, deleteQuestionRequest *question.DeleteQuestionRequest) (*question.DeleteQuestionResponse, error) {
	return questionHandler.questionController.DeleteQuestion(ctx, deleteQuestionRequest)
}
