package controller

import (
	"context"
	"database/sql"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/memory"
	modpb "memory_core/internal/proto/module"
	"memory_core/internal/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MemoryController struct {
	db              *sql.DB
	memoryRepo      *repository.MemoryRepository
	userModuleMapRepo   *repository.UserModuleMapRepository
	moduleRepo 	*repository.ModuleRepository
	docRepo *repository.DocRepository
}

func NewMemoryController(db *sql.DB, memoryRepo *repository.MemoryRepository, userModuleMapRepo *repository.UserModuleMapRepository, moduleRepo *repository.ModuleRepository, docRepo *repository.DocRepository) *MemoryController {
	return &MemoryController{db: db, memoryRepo: memoryRepo, userModuleMapRepo: userModuleMapRepo, moduleRepo: moduleRepo, docRepo: docRepo}
}

// CreateMemory handles the creation of a new memory
func (c *MemoryController) CreateMemory(ctx context.Context, req *mpb.CreateMemoryRequest) (*mpb.CreateMemoryResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}
	if req.GetMemoryTitle() == "" {
		log.Println("Memory title is required")
		return nil, status.Error(codes.InvalidArgument, "Memory title is required")
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

	// Create memory
	err = c.memoryRepo.CreateMemory(ctx, tx, req.GetUserId(), req.GetDocId(), req.GetMemoryTitle(), req.GetMemoryContent())
	if err != nil {
		log.Printf("Failed to create memory: %v", err)
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &mpb.CreateMemoryResponse{}, nil
}

// GetMemoryById retrieves a memory by its ID
func (c *MemoryController) GetMemoryById(ctx context.Context, req *mpb.GetMemoryByIdRequest) (*mpb.GetMemoryByIdResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetMemoryId() == 0 {
		log.Println("Memory ID is required")
		return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
	}
	// need doc ID to check if user has access to the memory
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get viewing privileges: %v", err)
		return nil, err
	}


	// Retrieve memory
	memory, err := c.memoryRepo.GetMemoryById(ctx, c.db, req.GetMemoryId())
	if err != nil {
		log.Printf("Failed to get memory by ID: %v", err)
		return nil, err
	}
	return &mpb.GetMemoryByIdResponse{Memory: memory}, nil
}

// GetMemoriesByDocId retrieves memories associated with a document ID
func (c *MemoryController) GetMemoriesByDocId(ctx context.Context, req *mpb.GetMemoriesByDocIdRequest) (*mpb.GetMemoriesByDocIdResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if user has access to the document
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get viewing privileges: %v", err)
		return nil, err
	}

	// Fetch memories
	memories, err := c.memoryRepo.GetMemoriesByDocId(ctx, c.db, req.GetDocId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get memories by doc ID: %v", err)
		return nil, err
	}

	return &mpb.GetMemoriesByDocIdResponse{Memories: memories}, nil
}

// GetMemoriesByTitleSearch retrieves memories by title search
func (c *MemoryController) GetMemoriesByTitleSearch(ctx context.Context, req *mpb.GetMemoriesByMemoryTitleSearchRequest) (*mpb.GetMemoriesByMemoryTitleSearchResponse, error) {
	// Validate input
	if req.GetSearchQuery() == "" {
		log.Println("Search query is required")
		return nil, status.Error(codes.InvalidArgument, "Search query is required")
	}else if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	// get viewing privileges
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get viewing privileges: %v", err)
		return nil, err
	}

	// Fetch memories
	memories, err := c.memoryRepo.GetMemoriesByMemoryTitleSearch(ctx, c.db, req.GetDocId(), req.GetSearchQuery(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get memories by title search: %v", err)
		return nil, err
	}

	return &mpb.GetMemoriesByMemoryTitleSearchResponse{Memories: memories}, nil
}

// UpdateMemory updates the content of an existing memory
func (c *MemoryController) UpdateMemory(ctx context.Context, req *mpb.UpdateMemoryRequest) (*mpb.UpdateMemoryResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetMemoryId() == 0 {
		log.Println("Memory ID is required")
		return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
	}

	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get edit privileges: %v", err)
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

	// Update memory
	err = c.memoryRepo.UpdateMemory(ctx, tx, req.GetMemoryId(), req.GetMemoryTitle(), req.GetMemoryContent())
	if err != nil {
		log.Printf("Failed to update memory: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Memory updated successfully")
	return &mpb.UpdateMemoryResponse{}, nil
}

// DeleteMemory deletes a memory
func (c *MemoryController) DeleteMemory(ctx context.Context, req *mpb.DeleteMemoryRequest) (*mpb.DeleteMemoryResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetMemoryId() == 0 {
		log.Println("Memory ID is required")
		return nil, status.Error(codes.InvalidArgument, "Memory ID is required")
	}

	// check if user has edit privileges
	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get edit privileges: %v", err)
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

	// Delete memory
	err = c.memoryRepo.DeleteMemory(ctx, tx, req.GetMemoryId())
	if err != nil {
		log.Printf("Failed to delete memory: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Memory deleted successfully")
	return &mpb.DeleteMemoryResponse{}, nil
}

// helper function to check if the user has access to the memory
func (c *MemoryController) getViewingPrivileges(ctx context.Context, userId, moduleId uint64) error {
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
func (c *MemoryController) getEditPrivileges(ctx context.Context, userId, moduleId uint64) error {
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