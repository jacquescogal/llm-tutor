package controller

import (
	"context"
	"database/sql"
	"log"
	"memory_core/internal/proto/common"
	qpb "memory_core/internal/proto/question"
	"memory_core/internal/repository"
	modpb "memory_core/internal/proto/module"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QuestionController struct {
	db              *sql.DB
	questionRepo    *repository.QuestionRepository
	docRepo         *repository.DocRepository
	moduleRepo      *repository.ModuleRepository
	userModuleMapRepo   *repository.UserModuleMapRepository
}

// NewQuestionController initializes a new QuestionController
func NewQuestionController(db *sql.DB, questionRepo *repository.QuestionRepository, docRepo *repository.DocRepository, moduleRepo *repository.ModuleRepository, userModuleMapRepo *repository.UserModuleMapRepository) *QuestionController {
	return &QuestionController{db: db, questionRepo: questionRepo, docRepo: docRepo, moduleRepo: moduleRepo, userModuleMapRepo: userModuleMapRepo}
}

// CreateQuestion handles the creation of a new question
func (c *QuestionController) CreateQuestion(ctx context.Context, req *qpb.CreateQuestionRequest) (*qpb.CreateQuestionResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}
	if req.GetQuestionTitle() == "" {
		log.Println("Question title is required")
		return nil, status.Error(codes.InvalidArgument, "Question title is required")
	}
	// check edit privileges
	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
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

	// Create question
	err = c.questionRepo.CreateQuestion(ctx, tx,  req.GetDocId(), req.GetUserId(), req.GetQuestionTitle(), req.GetQuestionBlob(), common.QuestionType(req.GetQuestionType()))
	if err != nil {
		log.Printf("Failed to create question: %v", err)
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &qpb.CreateQuestionResponse{}, nil
}

// GetQuestionById retrieves a question by its ID
func (c *QuestionController) GetQuestionById(ctx context.Context, req *qpb.GetQuestionByIdRequest) (*qpb.GetQuestionByIdResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetQuestionId() == 0 {
		log.Println("Question ID is required")
		return nil, status.Error(codes.InvalidArgument, "Question ID is required")
	}
	// check if have doc ID to check if user has access to the memory
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if user has access to the memory
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	// Retrieve question
	question, err := c.questionRepo.GetQuestionById(ctx, c.db, req.GetQuestionId())
	if err != nil {
		log.Printf("Failed to get question by ID: %v", err)
		return nil, err
	}

	return &qpb.GetQuestionByIdResponse{Question: question}, nil
}

// GetQuestionsByDocId retrieves questions associated with a document ID
func (c *QuestionController) GetQuestionsByDocId(ctx context.Context, req *qpb.GetQuestionsByDocIdRequest) (*qpb.GetQuestionsByDocIdResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if user has access to the question
	err := c.getViewingPrivileges(ctx, req.GetUserId(), req.GetModuleId())
	if err != nil {
		return nil, err
	}

	// Fetch questions
	questions, err := c.questionRepo.GetQuestionsByDocId(ctx, c.db, req.GetDocId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get questions by doc ID: %v", err)
		return nil, err
	}

	return &qpb.GetQuestionsByDocIdResponse{Questions: questions}, nil
}

// GetQuestionsByTitleSearch retrieves questions by title search
func (c *QuestionController) GetQuestionsByQuestionTitleSearch(ctx context.Context, req *qpb.GetQuestionsByQuestionTitleSearchRequest) (*qpb.GetQuestionsByQuestionTitleSearchResponse, error) {
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

	// Fetch questions
	questions, err := c.questionRepo.GetQuestionsByQuestionTitleSearch(ctx, c.db, req.GetDocId(), req.GetSearchQuery(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get questions by title search: %v", err)
		return nil, err
	}

	return &qpb.GetQuestionsByQuestionTitleSearchResponse{Questions: questions}, nil
}


// UpdateQuestion updates an existing question
func (c *QuestionController) UpdateQuestion(ctx context.Context, req *qpb.UpdateQuestionRequest) (*qpb.UpdateQuestionResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	if req.GetQuestionId() == 0 {
		log.Println("Question ID is required")
		return nil, status.Error(codes.InvalidArgument, "Question ID is required")
	}
	if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if user has access to the question
	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
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

	// Update question
	err = c.questionRepo.UpdateQuestion(ctx, tx, req.GetQuestionId(), req.GetQuestionTitle(), req.GetQuestionBlob(), req.GetQuestionType())
	if err != nil {
		log.Printf("Failed to update question: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &qpb.UpdateQuestionResponse{}, nil
}

// DeleteQuestion deletes a question
func (c *QuestionController) DeleteQuestion(ctx context.Context, req *qpb.DeleteQuestionRequest) (*qpb.DeleteQuestionResponse, error) {
	// Validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetQuestionId() == 0 {
		log.Println("Question ID is required")
		return nil, status.Error(codes.InvalidArgument, "Question ID is required")
	} else if req.GetDocId() == 0 {
		log.Println("Document ID is required")
		return nil, status.Error(codes.InvalidArgument, "Document ID is required")
	}

	// check if user has access to the question
	err := c.getEditPrivileges(ctx, req.GetUserId(), req.GetModuleId())
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

	// Delete question
	err = c.questionRepo.DeleteQuestion(ctx, tx, req.GetQuestionId())
	if err != nil {
		log.Printf("Failed to delete question: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	return &qpb.DeleteQuestionResponse{}, nil
}



// helper function to check if the user has access to the memory
func (c *QuestionController) getViewingPrivileges(ctx context.Context, userId, moduleId uint64) error {
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



// helper function to check if the user has edit access to module that contains the doc and question
func (c *QuestionController) getEditPrivileges(ctx context.Context, userId, moduleId uint64) error {

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