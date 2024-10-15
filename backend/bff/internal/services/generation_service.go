package services

import (
	"bff/internal/proto/generation"
	"context"
)

type GenerationService struct {
	client generation.GenerationServiceClient
}

func NewGenerationService(client generation.GenerationServiceClient) *GenerationService {
	return &GenerationService{client: client}
}

func (s *GenerationService) CreateGeneration(ctx context.Context, req *generation.CreateGenerationRequest) (*generation.CreateGenerationResponse, error) {
	res, err := s.client.CreateGeneration(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}