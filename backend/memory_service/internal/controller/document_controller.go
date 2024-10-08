package controller

// import (
// 	"context"
// 	"log"
// 	mpb "memory_core/internal/proto/memory"
// 	"memory_core/internal/repository"
// )

// type DocController struct {
// 	docRepo *repository.DocRepository
// }

// func NewDocController(docRepo *repository.DocRepository) *DocController {
// 	return &DocController{docRepo: docRepo}
// }

// // CreateDoc handles the business logic for creating a document
// func (c *DocController) CreateDoc(ctx context.Context, req *mpb.CreateDocRequest) (*mpb.CreateDocResponse, error) {
// 	err := c.docRepo.CreateDoc(ctx, req)
// 	if err != nil {
// 		log.Printf("Failed to create document: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.CreateDocResponse{}, nil
// }

// // GetDocById handles the business logic for retrieving a document by ID
// func (c *DocController) GetDocById(ctx context.Context, req *mpb.GetDocByIdRequest) (*mpb.GetDocByIdResponse, error) {
// 	doc, err := c.docRepo.GetDocById(ctx, req.DocId)
// 	if err != nil {
// 		log.Printf("Failed to get document by ID: %v", err)
// 		return nil, err
// 	}
// 	return &mpb.GetDocByIdResponse{Doc: doc}, nil
// }
