package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	dpb "memory_core/internal/proto/document"
	modpb "memory_core/internal/proto/module"

	"memory_core/internal/repository"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DocController struct {
	db *sql.DB
	docRepo *repository.DocRepository
	moduleRepo *repository.ModuleRepository
	userModuleMapRepo *repository.UserModuleMapRepository
}

func NewDocController(db *sql.DB, docRepo *repository.DocRepository, moduleRepo *repository.ModuleRepository, userModuleMapRepo *repository.UserModuleMapRepository) *DocController{
	return &DocController{db: db, docRepo: docRepo, moduleRepo: moduleRepo, userModuleMapRepo: userModuleMapRepo}
}

// CreateDoc handles the business logic for creating a new document
func (c *DocController) CreateDoc(ctx context.Context, req *dpb.CreateDocRequest) (*dpb.CreateDocResponse, error) {
	fmt.Println("CreateDoc")
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	validatedDocName, err := parseAndValidateDocName(req.GetDocName())
	fmt.Println("validatedDocName", validatedDocName)
	if err != nil {
		log.Printf("Invalid document name: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check edit privileges
	err = c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	// Begin transaction
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	req.DocName = validatedDocName

	// Create the document
	err = c.docRepo.CreateDoc(ctx, tx,  req.GetModuleId(), validatedDocName, req.GetDocDescription(), req.GetDocSummary(), req.GetUploadStatus(), req.GetS3ObjectKey())
	if err != nil {
		log.Printf("Failed to create document: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Document created successfully")
	return &dpb.CreateDocResponse{}, nil
}

// GetDocById handles the business logic for retrieving a document by ID
func (c *DocController) GetDocById(ctx context.Context, req *dpb.GetDocByIdRequest) (*dpb.GetDocByIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if the user has access to the module
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	doc, err := c.docRepo.GetDocById(ctx, c.db, req.GetDocId())
	if err != nil {
		log.Printf("Failed to get document by ID: %v", err)
		return nil, err
	}

	return &dpb.GetDocByIdResponse{Doc: doc}, nil
}

// GetDocsByModuleId handles the business logic for retrieving documents by module ID
func (c *DocController) GetDocsByModuleId(ctx context.Context, req *dpb.GetDocsByModuleIdRequest) (*dpb.GetDocsByModuleIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	// check if the user has access to the module
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	docs, err := c.docRepo.GetDocsByModuleId(ctx, c.db, req.GetModuleId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get documents by module ID: %v", err)
		return nil, err
	}

	return &dpb.GetDocsByModuleIdResponse{Docs: docs}, nil
}

// GetDocsByNameSearch handles the business logic for searching documents by name
func (c *DocController) GetDocsByNameSearch(ctx context.Context, req *dpb.GetDocsByNameSearchRequest) (*dpb.GetDocsByNameSearchResponse, error) {
	if req.GetSearchQuery() == "" {
		log.Println("Search query is required")
		return nil, status.Error(codes.InvalidArgument, "Search query is required")
	}
	// check module id
	if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
	}
	// check if the user has access to the module
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	docs, err := c.docRepo.GetDocsByNameSearch(ctx, c.db, req.GetModuleId(), req.GetSearchQuery(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to search documents by name: %v", err)
		return nil, err
	}

	return &dpb.GetDocsByNameSearchResponse{Docs: docs}, nil
}

// UpdateDoc handles the business logic for updating a document
func (c *DocController) UpdateDoc(ctx context.Context, req *dpb.UpdateDocRequest) (*dpb.UpdateDocResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	validatedDocName, err := parseAndValidateDocName(req.GetDocName())
	if err != nil {
		log.Printf("Invalid document name: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check edit privileges
	err = c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	err = c.docRepo.UpdateDoc(ctx, tx, req.GetDocId(), validatedDocName, req.GetDocDescription(), req.GetDocSummary(), req.GetUploadStatus(), req.GetS3ObjectKey())
	if err != nil {
		log.Printf("Failed to update document: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Document updated successfully")
	return &dpb.UpdateDocResponse{}, nil
}

// DeleteDoc handles the business logic for deleting a document
func (c *DocController) DeleteDoc(ctx context.Context, req *dpb.DeleteDocRequest) (*dpb.DeleteDocResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}
	// check edit privileges
	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	err = c.docRepo.DeleteDoc(ctx, tx, req.GetDocId())
	if err != nil {
		log.Printf("Failed to delete document: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Document deleted successfully")
	return &dpb.DeleteDocResponse{}, nil
}

// Utility functions for input validation
func parseAndValidateDocName(docName string) (string, error) {
	docName = strings.TrimSpace(docName)
	if len(docName) < 3 {
		return "", status.Error(codes.InvalidArgument, "Document name must be at least 3 characters")
	} else if len(docName) > 255 {
		return "", status.Error(codes.InvalidArgument, "Document name must be less than 255 characters")
	}
	return docName, nil
}



// helper function to check if the user has access to the memory
func (c *DocController) getViewingPrivileges(ctx context.Context, userId, moduleId uint64) error {
	// start concurrent fetch using goroutines
	var module *modpb.FullModule

	var err error
	module, err = c.moduleRepo.GetModuleById(ctx, c.db, userId, moduleId)
	if err != nil {
		log.Printf("Failed to get module by ID: %v", err)
	}

	// check if the module is public or if the user has access to the module
	if module == nil {
		return status.Error(codes.NotFound, "Module not found")
	} else if !module.Module.GetIsPublic() && (module.GetUserModuleRole() == common.UserModuleRole_USER_MODULE_ROLE_UNDEFINED) {
		return status.Error(codes.PermissionDenied, "User does not have permission to view module")
	}
	return nil
}

// helper function to check if the user has access to the memory
func (c *DocController) getEditPrivileges(ctx context.Context, userId, moduleId uint64) error {
	var module *modpb.FullModule

	var err error
	module, err = c.moduleRepo.GetModuleById(ctx, c.db, userId, moduleId)
	if err != nil {
		log.Printf("Failed to get module by ID: %v", err)
	}

	// check if the module is public or if the user has access to the module
	if module == nil {
		return status.Error(codes.NotFound, "Module not found")
	} else if (module.GetUserModuleRole() == common.UserModuleRole_USER_MODULE_ROLE_UNDEFINED || module.GetUserModuleRole() == common.UserModuleRole_USER_MODULE_ROLE_VIEWER) {
		return status.Error(codes.PermissionDenied, "User does not have permission to view module")
	}
	return nil
}