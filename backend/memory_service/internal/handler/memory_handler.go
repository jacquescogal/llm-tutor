package handler

import (
	"context"
	"memory_core/internal/controller"
	"memory_core/internal/proto/memory"
)

// MemoryHandler represents the handler for managing memories
type MemoryHandler struct {
	// Memory Controller
	memoryController *controller.MemoryController
	memory.UnimplementedMemoryServiceServer
}

// NewMemoryHandler creates a new MemoryHandler
func NewMemoryHandler(memoryController *controller.MemoryController) *MemoryHandler {
	return &MemoryHandler{memoryController: memoryController}
}

// CreateMemory creates a new memory
func (memoryHandler *MemoryHandler) CreateMemory(ctx context.Context, createMemoryRequest *memory.CreateMemoryRequest) (*memory.CreateMemoryResponse, error) {
	return memoryHandler.memoryController.CreateMemory(ctx, createMemoryRequest)
}

// GetMemoryById retrieves a memory by memory_id
func (memoryHandler *MemoryHandler) GetMemoryById(ctx context.Context, getMemoryByIdRequest *memory.GetMemoryByIdRequest) (*memory.GetMemoryByIdResponse, error) {
	return memoryHandler.memoryController.GetMemoryById(ctx, getMemoryByIdRequest)
}

// GetMemoriesByDocId retrieves all memories associated with a specific doc_id
func (memoryHandler *MemoryHandler) GetMemoriesByDocId(ctx context.Context, getMemoriesByDocIdRequest *memory.GetMemoriesByDocIdRequest) (*memory.GetMemoriesByDocIdResponse, error) {
	return memoryHandler.memoryController.GetMemoriesByDocId(ctx, getMemoriesByDocIdRequest)
}

// GetMemoriesByMemoryTitleSearch retrieves memories by a title search
func (memoryHandler *MemoryHandler) GetMemoriesByMemoryTitleSearch(ctx context.Context, getMemoriesByMemoryTitleSearchRequest *memory.GetMemoriesByMemoryTitleSearchRequest) (*memory.GetMemoriesByMemoryTitleSearchResponse, error) {
	return memoryHandler.memoryController.GetMemoriesByTitleSearch(ctx, getMemoriesByMemoryTitleSearchRequest)
}

// UpdateMemory updates a memory
func (memoryHandler *MemoryHandler) UpdateMemory(ctx context.Context, updateMemoryRequest *memory.UpdateMemoryRequest) (*memory.UpdateMemoryResponse, error) {
	return memoryHandler.memoryController.UpdateMemory(ctx, updateMemoryRequest)
}

// DeleteMemory deletes a memory
func (memoryHandler *MemoryHandler) DeleteMemory(ctx context.Context, deleteMemoryRequest *memory.DeleteMemoryRequest) (*memory.DeleteMemoryResponse, error) {
	return memoryHandler.memoryController.DeleteMemory(ctx, deleteMemoryRequest)
}
