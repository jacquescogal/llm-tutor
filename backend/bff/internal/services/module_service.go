package services

import (
	"bff/internal/proto/module"
	"context"
)

type ModuleService struct {
	client module.ModuleServiceClient
}

func NewModuleService(client module.ModuleServiceClient) *ModuleService {
	return &ModuleService{client: client}
}

func (s *ModuleService) CreateModule(ctx context.Context, req *module.CreateModuleRequest) error {
	_, err := s.client.CreateModule(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *ModuleService) GetPublicModules(ctx context.Context, req *module.GetPublicModulesRequest) (*module.GetPublicModulesResponse, error) {
	resp, err := s.client.GetPublicModules(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *ModuleService) GetPrivateModulesByUserId(ctx context.Context, req *module.GetPrivateModulesByUserIdRequest) (*module.GetPrivateModulesByUserIdResponse, error) {
	resp, err := s.client.GetPrivateModulesByUserId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *ModuleService) GetFavouriteModulesByUserId(ctx context.Context, req *module.GetFavouriteModulesByUserIdRequest) (*module.GetFavouriteModulesByUserIdResponse, error) {
	resp, err := s.client.GetFavouriteModulesByUserId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return resp, nil
}

func (s *ModuleService) GetModuleById(ctx context.Context, req *module.GetModuleByIdRequest) (*module.GetModuleByIdResponse, error) {
	res, err := s.client.GetModuleById(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *ModuleService) GetModulesBySubjectId(ctx context.Context, req *module.GetModulesBySubjectIdRequest) (*module.GetModulesBySubjectIdResponse, error) {
	res, err := s.client.GetModulesBySubjectId(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *ModuleService) GetModulesByNameSearch(ctx context.Context, req *module.GetModulesByNameSearchRequest) (*module.GetModulesByNameSearchResponse, error) {
	res, err := s.client.GetModulesByNameSearch(ctx, req)
	if err != nil {
		return nil, HandleGRPCError(err)
	}
	return res, nil
}

func (s *ModuleService) UpdateModule(ctx context.Context, req *module.UpdateModuleRequest) error {
	_, err := s.client.UpdateModule(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}

func (s *ModuleService) DeleteModule(ctx context.Context, req *module.DeleteModuleRequest) error {
	_, err := s.client.DeleteModule(ctx, req)
	if err != nil {
		return HandleGRPCError(err)
	}
	return nil
}