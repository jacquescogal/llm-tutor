package handler

import (
	"memory_core/internal/controller"
	"memory_core/internal/proto/vector"
	"context"
)

// VectorHandler represents the handler for managing vectors
type VectorHandler struct {
	// Vector
	vectorController *controller.VectorController
	vector.UnimplementedVectorServiceServer
}

// NewVectorHandler creates a new VectorHandler
func NewVectorHandler(vectorController *controller.VectorController) *VectorHandler {
	return &VectorHandler{vectorController: vectorController}
}

// CreateMemoryVector creates a new memory vector
func (vectorHandler *VectorHandler) CreateMemoryVector(ctx context.Context, createMemoryVectorRequest *vector.CreateMemoryVectorRequest) (*vector.CreateMemoryVectorResponse, error) {
	return vectorHandler.vectorController.CreateMemoryVector(ctx, createMemoryVectorRequest)
}

// SearchMemoryVector retrieves all memory vectors by search query
func (vectorHandler *VectorHandler) SearchMemoryVector(ctx context.Context, searchMemoryVectorRequest *vector.SearchMemoryVectorRequest) (*vector.SearchMemoryVectorResponse, error) {
	return vectorHandler.vectorController.SearchMemoryVector(ctx, searchMemoryVectorRequest)
}