package services

import (
	"bff/internal/proto/document"
	"context"
)

type DocumentService struct {
	client document.DocServiceClient
}

func NewDocumentService(client document.DocServiceClient) *DocumentService {
	return &DocumentService{client: client}
}

func (s *DocumentService) CreateDocument(ctx context.Context, req *document.CreateDocRequest) error {
	_, err := s.client.CreateDoc(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *DocumentService) GetDocumentById(ctx context.Context, req *document.GetDocByIdRequest) (*document.GetDocByIdResponse, error) {
	res, err := s.client.GetDocById(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *DocumentService) GetDocumentsByModuleId(ctx context.Context, req *document.GetDocsByModuleIdRequest) (*document.GetDocsByModuleIdResponse, error) {
	res, err := s.client.GetDocsByModuleId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *DocumentService) GetDocumentsByNameSearch(ctx context.Context, req *document.GetDocsByNameSearchRequest) (*document.GetDocsByNameSearchResponse, error) {
	res, err := s.client.GetDocsByNameSearch(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *DocumentService) UpdateDocument(ctx context.Context, req *document.UpdateDocRequest) error {
	_, err := s.client.UpdateDoc(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *DocumentService) DeleteDocument(ctx context.Context, req *document.DeleteDocRequest) error {
	_, err := s.client.DeleteDoc(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}