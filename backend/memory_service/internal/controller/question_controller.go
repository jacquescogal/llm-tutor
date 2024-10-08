package controller

// import (
// 	"context"
// 	"log"
// 	mpb "memory_core/internal/proto/memory"
// 	"memory_core/internal/repository"
// )

// type QuestionController struct {
// 	questionRepo *repository.QuestionRepository
// }

// func NewQuestionController(questionRepo *repository.QuestionRepository) *QuestionController {
// 	return &QuestionController{questionRepo: questionRepo}
// }

// // CreateQuestion handles the business logic for creating a question
// func (c *QuestionController) CreateQuestion(ctx context.Context, req *mpb.CreateQuestionRequest) (*mpb.CreateQuestionResponse, error) {
// 	err := c.questionRepo.CreateQuestion(ctx, req)
// 	if err != nil {
// 		log.Printf("Failed to create question: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.CreateQuestionResponse{}, nil
// }

// // GetQuestionById handles the business logic for retrieving a question by ID
// func (c *QuestionController) GetQuestionById(ctx context.Context, req *mpb.GetQuestionByIdRequest) (*mpb.GetQuestionByIdResponse, error) {
// 	question, err := c.questionRepo.GetQuestionById(ctx, req.QuestionId)
// 	if err != nil {
// 		log.Printf("Failed to get question by ID: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.GetQuestionByIdResponse{Question: question}, nil
// }
