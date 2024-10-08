package handler

import (
	"context"
	"memory_core/internal/controller"
	"memory_core/internal/proto/document"
)

// DocHandler represents the handler for managing documents
type DocHandler struct {
	// Doc Controller
	docController *controller.DocController
	document.UnimplementedDocServiceServer
}

// NewDocHandler creates a new DocHandler
func NewDocHandler(docController *controller.DocController) *DocHandler {
	return &DocHandler{docController: docController}
}

// CreateDoc creates a new document
func (docHandler *DocHandler) CreateDoc(ctx context.Context, createDocRequest *document.CreateDocRequest) (*document.CreateDocResponse, error) {
	return docHandler.docController.CreateDoc(ctx, createDocRequest)
}

// GetDocById retrieves a document by its doc_id
func (docHandler *DocHandler) GetDocById(ctx context.Context, getDocByIdRequest *document.GetDocByIdRequest) (*document.GetDocByIdResponse, error) {
	return docHandler.docController.GetDocById(ctx, getDocByIdRequest)
}

// GetDocsByModuleId retrieves all documents associated with a specific module_id
func (docHandler *DocHandler) GetDocsByModuleId(ctx context.Context, getDocsByModuleIdRequest *document.GetDocsByModuleIdRequest) (*document.GetDocsByModuleIdResponse, error) {
	return docHandler.docController.GetDocsByModuleId(ctx, getDocsByModuleIdRequest)
}

// GetDocsByNameSearch retrieves all documents by name search
func (docHandler *DocHandler) GetDocsByNameSearch(ctx context.Context, getDocsByNameSearchRequest *document.GetDocsByNameSearchRequest) (*document.GetDocsByNameSearchResponse, error) {
	return docHandler.docController.GetDocsByNameSearch(ctx, getDocsByNameSearchRequest)
}

// UpdateDoc updates a document
func (docHandler *DocHandler) UpdateDoc(ctx context.Context, updateDocRequest *document.UpdateDocRequest) (*document.UpdateDocResponse, error) {
	return docHandler.docController.UpdateDoc(ctx, updateDocRequest)
}

// DeleteDoc deletes a document
func (docHandler *DocHandler) DeleteDoc(ctx context.Context, deleteDocRequest *document.DeleteDocRequest) (*document.DeleteDocResponse, error) {
	return docHandler.docController.DeleteDoc(ctx, deleteDocRequest)
}
