package controller

// import (
// 	"context"
// 	"log"
// 	mpb "memory_core/internal/proto/memory"
// 	"memory_core/internal/repository"
// )

// type MemoryController struct {
// 	memoryRepo *repository.MemoryRepository
// }

// func NewMemoryController(memoryRepo *repository.MemoryRepository) *MemoryController {
// 	return &MemoryController{memoryRepo: memoryRepo}
// }

// // CreateMemory handles the business logic for creating a memory
// func (c *MemoryController) CreateMemory(ctx context.Context, req *mpb.CreateMemoryRequest) (*mpb.CreateMemoryResponse, error) {
// 	err := c.memoryRepo.CreateMemory(ctx, req)
// 	if err != nil {
// 		log.Printf("Failed to create memory: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.CreateMemoryResponse{}, nil
// }

// // GetMemoryById handles the business logic for retrieving a memory by ID
// func (c *MemoryController) GetMemoryById(ctx context.Context, req *mpb.GetMemoryByIdRequest) (*mpb.GetMemoryByIdResponse, error) {
// 	memory, err := c.memoryRepo.GetMemoryById(ctx, req.MemoryId)
// 	if err != nil {
// 		log.Printf("Failed to get memory by ID: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.GetMemoryByIdResponse{Memory: memory}, nil
// }
