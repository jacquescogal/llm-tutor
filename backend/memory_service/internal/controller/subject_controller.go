package controller

import (
	"context"
	"database/sql"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/subject"
	"memory_core/internal/repository"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SubjectController struct {
	db *sql.DB
	subjectRepo *repository.SubjectRepository
	userSubjectMapRepo *repository.UserSubjectMapRepository
}

func NewSubjectController(db *sql.DB, subjectRepo *repository.SubjectRepository, userSubjectMapRepo *repository.UserSubjectMapRepository) *SubjectController {
	return &SubjectController{db: db, subjectRepo: subjectRepo, userSubjectMapRepo: userSubjectMapRepo}
}

// CreateSubject handles the business logic for creating a new subject
func (c *SubjectController) CreateSubject(ctx context.Context, req *mpb.CreateSubjectRequest) (*mpb.CreateSubjectResponse, error) {
	// create subject and user_id - subject_id admin role mapping to member_access_tab
	// atomic transaction

	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} 
	validatedSubjectName, err := parseAndValidateSubjectName(req.GetSubjectName())
	if err != nil {
		log.Printf("Invalid subject name: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}


	// begin transaction
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

	// create the subject
	subjectID, err := c.subjectRepo.CreateSubject(ctx, tx, validatedSubjectName, req.GetSubjectDescription(), req.GetIsPublic())
	if err != nil {
		log.Printf("Failed to create subject: %v", err)
		tx.Rollback()
		return nil, err
	}

	// create the member access mapping
	err = c.userSubjectMapRepo.PutUserSubjectMapping(ctx, tx, req.GetUserId(), subjectID, common.UserSubjectRole_USER_SUBJECT_ROLE_OWNER, false)
	if err != nil {
		log.Printf("Failed to create member access: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Subject created successfully")
	return &mpb.CreateSubjectResponse{}, nil
}

// GetSubjectById handles the business logic for retrieving a subject by ID
func (c *SubjectController) GetSubjectById(ctx context.Context, req *mpb.GetSubjectByIdRequest) (*mpb.GetSubjectByIdResponse, error) {
	// validate input
	if (req.GetUserId() == 0) {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if (req.GetSubjectId() == 0) {
		log.Println("Subject ID is required")
		return nil, status.Error(codes.InvalidArgument, "Subject ID is required")
	}
	subject, err := c.subjectRepo.GetSubjectById(ctx, c.db, req.GetUserId(), req.GetSubjectId())
	if err != nil {
		log.Printf("Failed to get subject by ID: %v", err)
		return nil, err
	}
	return &mpb.GetSubjectByIdResponse{Subject: subject}, nil
}

// GetSubjectsByUserId handles the business logic for retrieving subjects by user ID
func (c *SubjectController) GetSubjectsByUserId(ctx context.Context, req *mpb.GetSubjectsByUserIdRequest) (*mpb.GetSubjectsByUserIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	subjects, err := c.subjectRepo.GetSubjectsByUserId(ctx, c.db, req.GetUserId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get subjects by user ID: %v", err)
		return nil, err
	}
	return &mpb.GetSubjectsByUserIdResponse{Subjects: subjects}, nil
}

// GetSubjectsByNameSearch handles the business logic for retrieving subjects by name search
func (c *SubjectController) GetSubjectsByNameSearch(ctx context.Context, req *mpb.GetSubjectsByNameSearchRequest) (*mpb.GetSubjectsByNameSearchResponse, error) {
	if req.GetSearchQuery() == "" {
		log.Println("Search Query is required")
		return nil, status.Error(codes.InvalidArgument, "Name search is required")
	} else if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	subjects, err := c.subjectRepo.GetSubjectsByNameSearch(ctx, c.db, req.GetUserId(), req.GetSearchQuery(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get subjects by name search: %v", err)
		return nil, err
	}
	return &mpb.GetSubjectsByNameSearchResponse{Subjects: subjects}, nil
}

// UpdateSubject handles the business logic for updating a subject
func (c *SubjectController) UpdateSubject(ctx context.Context, req *mpb.UpdateSubjectRequest) (*mpb.UpdateSubjectResponse, error) {
	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetSubjectId() == 0 {
		log.Println("Subject ID is required")
		return nil, status.Error(codes.InvalidArgument, "Subject ID is required")
	} 
	validatedSubjectName, err := parseAndValidateSubjectName(req.GetSubjectName())
	if err != nil {
		log.Printf("Invalid subject name: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check if user is an admin, owner or editor of the subject
	memberAccess, err := c.userSubjectMapRepo.GetUserSubjectMapping(ctx, c.db, req.GetUserId(), req.GetSubjectId())
	if err != nil {
		log.Printf("Failed to get member access: %v", err)
		return nil, err
	} else if err := validateMemberEditPrivileges(memberAccess.GetUserSubjectRole()); err != nil {
		log.Printf("User does not have permission to edit subject: %v", err)
		return nil, err
	}

	// begin transaction
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
	// TODO: lock this subject row for update in distributed lock
	// update the subject
	err = c.subjectRepo.UpdateSubject(ctx, tx, req.GetSubjectId(), validatedSubjectName, req.GetSubjectDescription(), req.GetIsPublic())
	if err != nil {
		log.Printf("Failed to update subject: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Subject updated successfully")
	return &mpb.UpdateSubjectResponse{}, nil
}

// DeleteSubject handles the business logic for deleting a subject
func (c *SubjectController) DeleteSubject(ctx context.Context, req *mpb.DeleteSubjectRequest) (*mpb.DeleteSubjectResponse, error) {
	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetSubjectId() == 0 {
		log.Println("Subject ID is required")
		return nil, status.Error(codes.InvalidArgument, "Subject ID is required")
	}

	// check if user is an admin or owner of the subject
	memberAccess, err := c.userSubjectMapRepo.GetUserSubjectMapping(ctx, c.db, req.GetUserId(), req.GetSubjectId())
	if err != nil {
		log.Printf("Failed to get member access: %v", err)
		return nil, err
	} else if err := validateMemberEditPrivileges(memberAccess.GetUserSubjectRole()); err != nil {
		log.Printf("User does not have permission to edit subject: %v", err)
		return nil, err
	}

	// begin transaction
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

	// delete the subject
	err = c.subjectRepo.DeleteSubject(ctx, tx, req.GetSubjectId())
	if err != nil {
		log.Printf("Failed to delete subject: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Subject deleted successfully")
	return &mpb.DeleteSubjectResponse{}, nil
}

func parseAndValidateSubjectName(subjectName string) (string, error) {
	// strip leading and trailing whitespace
	subjectName = strings.TrimSpace(subjectName)
	if len(subjectName) < 3 {
		return "", status.Error(codes.InvalidArgument, "Subject name must be at least 3 characters")
	} else if len(subjectName) > 100 {
		return "", status.Error(codes.InvalidArgument, "Subject name must be less than 100 characters")
	} else{
		return subjectName, nil
	}
}

func validateMemberEditPrivileges(memberRole common.UserSubjectRole) error {
	if memberRole != common.UserSubjectRole_USER_SUBJECT_ROLE_ADMIN && memberRole != common.UserSubjectRole_USER_SUBJECT_ROLE_OWNER && memberRole != common.UserSubjectRole_USER_SUBJECT_ROLE_EDITOR {
		return status.Error(codes.PermissionDenied, "User does not have permission to edit subject")
	}
	return nil
}