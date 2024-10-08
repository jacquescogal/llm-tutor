package handler

import (
	"memory_core/internal/controller"
	"memory_core/internal/proto/subject"
	"context"
)

// SubjectHandler represents the handler for managing subjects
type SubjectHandler struct {
	// Subject
	subjectController *controller.SubjectController
	subject.UnimplementedSubjectServiceServer
}

// NewSubjectHandler creates a new SubjectHandler
func NewSubjectHandler(subjectController *controller.SubjectController) *SubjectHandler {
	return &SubjectHandler{subjectController: subjectController}
}

// CreateSubject creates a new subject
func (subjectHandler *SubjectHandler) CreateSubject(ctx context.Context, createSubjectRequest *subject.CreateSubjectRequest) (*subject.CreateSubjectResponse, error) {
	return subjectHandler.subjectController.CreateSubject(ctx, createSubjectRequest)
}

// GetPublicSubjects retrieves all public subjects
func (subjectHandler *SubjectHandler) GetPublicSubjects(ctx context.Context, getPublicSubjectsRequest *subject.GetPublicSubjectsRequest) (*subject.GetPublicSubjectsResponse, error) {
	return subjectHandler.subjectController.GetPublicSubjects(ctx, getPublicSubjectsRequest)
}

// GetPrivateSubjectsByUserId retrieves all private subjects by user_id
func (subjectHandler *SubjectHandler) GetPrivateSubjectsByUserId(ctx context.Context, getPrivateSubjectsByUserIdRequest *subject.GetPrivateSubjectsByUserIdRequest) (*subject.GetPrivateSubjectsByUserIdResponse, error) {
	return subjectHandler.subjectController.GetPrivateSubjectsByUserId(ctx, getPrivateSubjectsByUserIdRequest)
}

// GetFavouriteSubjectsByUserId retrieves all favourite subjects by user_id
func (subjectHandler *SubjectHandler) GetFavouriteSubjectsByUserId(ctx context.Context, getFavouriteSubjectsByUserIdRequest *subject.GetFavouriteSubjectsByUserIdRequest) (*subject.GetFavouriteSubjectsByUserIdResponse, error) {
	return subjectHandler.subjectController.GetFavouriteSubjectsByUserId(ctx, getFavouriteSubjectsByUserIdRequest)
}

// GetSubjectById retrieves a subject by subject_id
func (subjectHandler *SubjectHandler) GetSubjectById(ctx context.Context, getSubjectByIdRequest *subject.GetSubjectByIdRequest) (*subject.GetSubjectByIdResponse, error) {
	return subjectHandler.subjectController.GetSubjectById(ctx, getSubjectByIdRequest)
}

// GetSubjectsByUserId retrieves all subjects by user_id
func (subjectHandler *SubjectHandler) GetSubjectsByUserId(ctx context.Context, getSubjectsByUserIdRequest *subject.GetSubjectsByUserIdRequest) (*subject.GetSubjectsByUserIdResponse, error) {
	return subjectHandler.subjectController.GetSubjectsByUserId(ctx, getSubjectsByUserIdRequest)
}

// GetSubjectsByNameSearch retrieves all subjects by name search
func (subjectHandler *SubjectHandler) GetSubjectsByNameSearch(ctx context.Context, getSubjectsByNameSearchRequest *subject.GetSubjectsByNameSearchRequest) (*subject.GetSubjectsByNameSearchResponse, error) {
	return subjectHandler.subjectController.GetSubjectsByNameSearch(ctx, getSubjectsByNameSearchRequest)
}

// UpdateSubject updates a subject
func (subjectHandler *SubjectHandler) UpdateSubject(ctx context.Context, updateSubjectRequest *subject.UpdateSubjectRequest) (*subject.UpdateSubjectResponse, error) {
	return subjectHandler.subjectController.UpdateSubject(ctx, updateSubjectRequest)
}

// DeleteSubject deletes a subject
func (subjectHandler *SubjectHandler) DeleteSubject(ctx context.Context, deleteSubjectRequest *subject.DeleteSubjectRequest) (*subject.DeleteSubjectResponse, error) {
	return subjectHandler.subjectController.DeleteSubject(ctx, deleteSubjectRequest)
}

func (subjectHandler *SubjectHandler) SetUserSubjectFavourite(ctx context.Context, setUserSubjectFavouriteRequest *subject.SetUserSubjectFavouriteRequest) (*subject.SetUserSubjectFavouriteResponse, error) {
	return subjectHandler.subjectController.SetUserSubjectFavourite(ctx, setUserSubjectFavouriteRequest)
}