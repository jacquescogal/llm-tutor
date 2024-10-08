package handler

import (
	"memory_core/internal/controller"
	"memory_core/internal/proto/module"
	"context"
)

// ModuleHandler represents the handler for managing modules
type ModuleHandler struct {
	// Module
	moduleController *controller.ModuleController
	module.UnimplementedModuleServiceServer
}

// NewModuleHandler creates a new ModuleHandler
func NewModuleHandler(moduleController *controller.ModuleController) *ModuleHandler {
	return &ModuleHandler{moduleController: moduleController}
}

// CreateModule creates a new module
func (moduleHandler *ModuleHandler) CreateModule(ctx context.Context, createModuleRequest *module.CreateModuleRequest) (*module.CreateModuleResponse, error) {
	return moduleHandler.moduleController.CreateModule(ctx, createModuleRequest)
}

// GetPublicModules retrieves all public modules
func (moduleHandler *ModuleHandler) GetPublicModules(ctx context.Context, getPublicModulesRequest *module.GetPublicModulesRequest) (*module.GetPublicModulesResponse, error) {
	return moduleHandler.moduleController.GetPublicModules(ctx, getPublicModulesRequest)
}

// GetPrivateModulesByUserId retrieves all private modules by user_id
func (moduleHandler *ModuleHandler) GetPrivateModulesByUserId(ctx context.Context, getPrivateModulesByUserIdRequest *module.GetPrivateModulesByUserIdRequest) (*module.GetPrivateModulesByUserIdResponse, error) {
	return moduleHandler.moduleController.GetPrivateModulesByUserId(ctx, getPrivateModulesByUserIdRequest)
}

// GetFavouriteModulesByUserId retrieves all favourite modules by user_id
func (moduleHandler *ModuleHandler) GetFavouriteModulesByUserId(ctx context.Context, getFavouriteModulesByUserIdRequest *module.GetFavouriteModulesByUserIdRequest) (*module.GetFavouriteModulesByUserIdResponse, error) {
	return moduleHandler.moduleController.GetFavouriteModulesByUserId(ctx, getFavouriteModulesByUserIdRequest)
}
// GetModuleById retrieves a module by module_id
func (moduleHandler *ModuleHandler) GetModuleById(ctx context.Context, getModuleByIdRequest *module.GetModuleByIdRequest) (*module.GetModuleByIdResponse, error) {
	return moduleHandler.moduleController.GetModuleById(ctx, getModuleByIdRequest)
}

// GetModulesBySubjectId retrieves all modules by subject_id
func (moduleHandler *ModuleHandler) GetModulesBySubjectId(ctx context.Context, getModulesBySubjectIdRequest *module.GetModulesBySubjectIdRequest) (*module.GetModulesBySubjectIdResponse, error) {
	return moduleHandler.moduleController.GetModulesBySubjectId(ctx, getModulesBySubjectIdRequest)
}

// GetModulesByNameSearch retrieves all modules by name search
func (moduleHandler *ModuleHandler) GetModulesByNameSearch(ctx context.Context, getModulesByNameSearchRequest *module.GetModulesByNameSearchRequest) (*module.GetModulesByNameSearchResponse, error) {
	return moduleHandler.moduleController.GetModulesByNameSearch(ctx, getModulesByNameSearchRequest)
}


// UpdateModule updates a module
func (moduleHandler *ModuleHandler) UpdateModule(ctx context.Context, updateModuleRequest *module.UpdateModuleRequest) (*module.UpdateModuleResponse, error) {
	return moduleHandler.moduleController.UpdateModule(ctx, updateModuleRequest)
}

// DeleteModule deletes a module
func (moduleHandler *ModuleHandler) DeleteModule(ctx context.Context, deleteModuleRequest *module.DeleteModuleRequest) (*module.DeleteModuleResponse, error) {
	return moduleHandler.moduleController.DeleteModule(ctx, deleteModuleRequest)
}
