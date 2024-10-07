package services

import (
	"bff/internal/proto/question"
	"context"
)

type QuestionService struct {
	client question.QuestionServiceClient
}

func NewQuestionService(client question.QuestionServiceClient) *QuestionService {
	return &QuestionService{client: client}
}

func (s *QuestionService) CreateQuestion(ctx context.Context, req *question.CreateQuestionRequest) error {
	_, err := s.client.CreateQuestion(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *QuestionService) GetQuestionById(ctx context.Context, req *question.GetQuestionByIdRequest) (*question.GetQuestionByIdResponse, error) {
	res,err := s.client.GetQuestionById(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *QuestionService) GetQuestionsByDocId(ctx context.Context, req *question.GetQuestionsByDocIdRequest) (*question.GetQuestionsByDocIdResponse, error) {
	res,err := s.client.GetQuestionsByDocId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *QuestionService) GetQuestionsByQuestionTitleSearch(ctx context.Context,req *question.GetQuestionsByQuestionTitleSearchRequest) (*question.GetQuestionsByQuestionTitleSearchResponse, error) {
	res,err := s.client.GetQuestionsByQuestionTitleSearch(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *QuestionService) UpdateQuestion(ctx context.Context, req *question.UpdateQuestionRequest) error {
	_, err := s.client.UpdateQuestion(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *QuestionService) DeleteQuestion(ctx context.Context, req *question.DeleteQuestionRequest) error {
	_, err := s.client.DeleteQuestion(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}