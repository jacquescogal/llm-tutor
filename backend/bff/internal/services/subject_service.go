package services

import (
	"bff/internal/proto/subject"
	"context"
)

type SubjectService struct {
	client subject.SubjectServiceClient
}

func NewSubjectService(client subject.SubjectServiceClient) *SubjectService {
	return &SubjectService{client: client}
}

func (s *SubjectService) CreateSubject(ctx context.Context, req *subject.CreateSubjectRequest) error {
	_, err := s.client.CreateSubject(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *SubjectService) GetPublicSubjects(ctx context.Context, req *subject.GetPublicSubjectsRequest) (*subject.GetPublicSubjectsResponse, error) {
	resp, err := s.client.GetPublicSubjects(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *SubjectService) GetPrivateSubjectsByUserID(ctx context.Context, req *subject.GetPrivateSubjectsByUserIdRequest) (*subject.GetPrivateSubjectsByUserIdResponse, error) {
	resp, err := s.client.GetPrivateSubjectsByUserId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *SubjectService) GetFavouriteSubjectsByUserID(ctx context.Context, req *subject.GetFavouriteSubjectsByUserIdRequest) (*subject.GetFavouriteSubjectsByUserIdResponse, error) {
	resp, err := s.client.GetFavouriteSubjectsByUserId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *SubjectService) GetSubjectByID(ctx context.Context, req *subject.GetSubjectByIdRequest) (*subject.GetSubjectByIdResponse, error) {
	resp, err := s.client.GetSubjectById(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *SubjectService) GetSubjectsByUserID(ctx context.Context, req *subject.GetSubjectsByUserIdRequest) (*subject.GetSubjectsByUserIdResponse, error) {
	resp, err := s.client.GetSubjectsByUserId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *SubjectService) GetSubjectsByNameSearch(ctx context.Context, req *subject.GetSubjectsByNameSearchRequest) (*subject.GetSubjectsByNameSearchResponse, error) {
	res,  err := s.client.GetSubjectsByNameSearch(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *SubjectService) UpdateSubject(ctx context.Context, req *subject.UpdateSubjectRequest) error {
	_, err := s.client.UpdateSubject(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *SubjectService) DeleteSubject(ctx context.Context, req *subject.DeleteSubjectRequest) error {
	_, err := s.client.DeleteSubject(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}