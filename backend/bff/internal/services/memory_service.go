package services

import (
	"bff/internal/proto/memory"
	"context"
)

type MemoryService struct {
    client memory.MemoryServiceClient
}

func NewMemoryService(client memory.MemoryServiceClient) *MemoryService {
    return &MemoryService{client: client}
}

func (s *MemoryService) CreateMemory(ctx context.Context, req *memory.CreateMemoryRequest) error {
    _, err := s.client.CreateMemory(ctx, req)
    if err != nil {
        return HandleGRPCError(err)
    }
    return nil
}

func (s *MemoryService) GetMemoryById(ctx context.Context, req *memory.GetMemoryByIdRequest) (*memory.GetMemoryByIdResponse, error) {
    res, err := s.client.GetMemoryById(ctx, req)
    if err != nil {
        return nil, HandleGRPCError(err)
    }
    return res, nil
}

func (s *MemoryService) GetMemoriesByDocId(ctx context.Context, req *memory.GetMemoriesByDocIdRequest) (*memory.GetMemoriesByDocIdResponse, error) {
    res, err := s.client.GetMemoriesByDocId(ctx, req)
    if err != nil {
        return nil, HandleGRPCError(err)
    }
    return res, nil
}

func (s *MemoryService) GetMemoriesByMemoryTitleSearch(ctx context.Context, req *memory.GetMemoriesByMemoryTitleSearchRequest) (*memory.GetMemoriesByMemoryTitleSearchResponse, error) {
    res, err := s.client.GetMemoriesByMemoryTitleSearch(ctx, req)
    if err != nil {
        return nil, HandleGRPCError(err)
    }
    return res, nil
}

func (s *MemoryService) UpdateMemory(ctx context.Context,req *memory.UpdateMemoryRequest) error {
    _, err := s.client.UpdateMemory(ctx, req)
    if err != nil {
        return HandleGRPCError(err)
    }
    return nil
}

func (s *MemoryService) DeleteMemory(ctx context.Context, req *memory.DeleteMemoryRequest) error {
    _, err := s.client.DeleteMemory(ctx, req)
    if err != nil {
        return HandleGRPCError(err)
    }
    return nil
}