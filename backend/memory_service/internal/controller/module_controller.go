package controller

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"memory_core/internal/proto/common"
	mpb "memory_core/internal/proto/module"
	"memory_core/internal/repository"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ModuleController struct {
	db *sql.DB
	moduleRepo *repository.ModuleRepository
	userModuleMapRepo *repository.UserModuleMapRepository
}

func NewModuleController(db *sql.DB, moduleRepo *repository.ModuleRepository, userModuleMapRepo *repository.UserModuleMapRepository) *ModuleController{
	return &ModuleController{db: db, moduleRepo: moduleRepo, userModuleMapRepo: userModuleMapRepo}
}

func (c *ModuleController) CreateModule(ctx context.Context, req *mpb.CreateModuleRequest) (*mpb.CreateModuleResponse, error) {
	// create module and user_id - module_id admin role mapping to member_access_tab
	// atomic transaction

	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} 
	validatedModuleName, err := parseAndValidateModuleName(req.GetModuleName())
	if err != nil {
		log.Printf("Invalid module name: %v", err)
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

	// create the module
	moduleId, err := c.moduleRepo.CreateModule(ctx, tx, validatedModuleName, req.GetModuleDescription(), req.GetIsPublic())
	if err != nil {
		log.Printf("Failed to create module: %v", err)
		return nil, err
	}

	// create the mapping
	err = c.userModuleMapRepo.PutUserModuleMappingRole(ctx, tx, req.GetUserId(), moduleId, common.UserModuleRole_USER_MODULE_ROLE_OWNER)
	if err != nil {
		log.Printf("Failed to create user module mapping: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Module created successfully")
	return &mpb.CreateModuleResponse{}, nil
}

// GetPublicModules handles the business logic for retrieving all public modules
func (c *ModuleController) GetPublicModules(ctx context.Context, req *mpb.GetPublicModulesRequest) (*mpb.GetPublicModulesResponse, error) {
	modules, err := c.moduleRepo.GetPublicModules(ctx, c.db, req.UserId, req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get public modules: %v", err)
		return nil, err
	}
	return &mpb.GetPublicModulesResponse{Modules: modules}, nil
}

// GetPrivateModulesByUserId handles the business logic for retrieving private modules by user ID
func (c *ModuleController) GetPrivateModulesByUserId(ctx context.Context, req *mpb.GetPrivateModulesByUserIdRequest) (*mpb.GetPrivateModulesByUserIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	modules, err := c.moduleRepo.GetPrivateModulesByUserId(ctx, c.db, req.GetUserId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get private modules by user ID: %v", err)
		return nil, err
	}
	return &mpb.GetPrivateModulesByUserIdResponse{Modules: modules}, nil
}

// GetFavouriteModulesByUserId handles the business logic for retrieving favorite modules by user ID
func (c *ModuleController) GetFavouriteModulesByUserId(ctx context.Context, req *mpb.GetFavouriteModulesByUserIdRequest) (*mpb.GetFavouriteModulesByUserIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	modules, err := c.moduleRepo.GetFavouriteModulesByUserId(ctx, c.db, req.GetUserId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get favorite modules by user ID: %v", err)
		return nil, err
	}
	return &mpb.GetFavouriteModulesByUserIdResponse{Modules: modules}, nil
}

func (c *ModuleController) GetModuleById(ctx context.Context, req *mpb.GetModuleByIdRequest) (*mpb.GetModuleByIdResponse, error) {
	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
	}
	// start concurrent fetch using goroutines
	var module *mpb.FullModule

	var err error
	module, err = c.moduleRepo.GetModuleById(ctx, c.db, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get module by ID: %v", err)
	}
	

	// check if the module is public or if the user has access to the module
	if module == nil {
		return nil, status.Error(codes.NotFound, "Module not found")
	} else if !module.Module.GetIsPublic() && (module.GetUserModuleRole() == common.UserModuleRole_USER_MODULE_ROLE_UNDEFINED) {
		return nil, status.Error(codes.PermissionDenied, "User does not have permission to view module")
	}
	return &mpb.GetModuleByIdResponse{Module: module}, nil
}

func (c *ModuleController) GetModulesBySubjectId(ctx context.Context, req *mpb.GetModulesBySubjectIdRequest) (*mpb.GetModulesBySubjectIdResponse, error) {
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}

	modules, err := c.moduleRepo.GetModulesBySubjectId(ctx, c.db, req.GetUserId(), req.GetSubjectId(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get modules by subject ID: %v", err)
		return nil, err
	}
	fmt.Println(modules)
	return &mpb.GetModulesBySubjectIdResponse{Modules: modules}, nil
}

func (c *ModuleController) GetModulesByNameSearch(ctx context.Context, req *mpb.GetModulesByNameSearchRequest) (*mpb.GetModulesByNameSearchResponse, error) {
	if req.GetSearchQuery() == "" {
		log.Println("Search Query is required")
		return nil, status.Error(codes.InvalidArgument, "Name search is required")
	} else if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	}
	// todo: decide if we want to allow public modules to be searched by anyone
	modules, err := c.moduleRepo.GetPublicModulesByNameSearch(ctx, c.db, req.GetUserId(), req.GetSearchQuery(), req.GetPageNumber(), req.GetPageSize(), req.GetOrderByField(), req.GetOrderByDirection())
	if err != nil {
		log.Printf("Failed to get modules by name search: %v", err)
		return nil, err
	}
	return &mpb.GetModulesByNameSearchResponse{Modules: modules}, nil
}

func (c *ModuleController) UpdateModule(ctx context.Context, req *mpb.UpdateModuleRequest) (*mpb.UpdateModuleResponse, error) {

	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
	} 
	validatedModuleName, err := parseAndValidateModuleName(req.GetModuleName())
	if err != nil {
		log.Printf("Invalid module name: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// check if user is an admin, owner or editor of the module
	userModuleAccessRole, err := c.userModuleMapRepo.GetUserModuleMapping(ctx, c.db, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get user module access: %v", err)
		return nil, err
	} else if err := validateUserModuleAccess(userModuleAccessRole.GetUserModuleRole()); err != nil {
		log.Printf("User does not have permission to edit module: %v", err)
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

	// update the module
	err = c.moduleRepo.UpdateModule(ctx, tx, req.GetModuleId(), validatedModuleName, req.GetModuleDescription(), req.GetIsPublic())
	if err != nil {
		log.Printf("Failed to update module: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Module updated successfully")
	return &mpb.UpdateModuleResponse{}, nil
}

func (c *ModuleController) DeleteModule(ctx context.Context, req *mpb.DeleteModuleRequest) (*mpb.DeleteModuleResponse, error) {
	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
	}

	// check if user is an admin or owner of the module
	userModuleAccessRole, err := c.userModuleMapRepo.GetUserModuleMapping(ctx, c.db, req.GetUserId(), req.GetModuleId())
	if err != nil {
		log.Printf("Failed to get user module access: %v", err)
		return nil, err
	} else if err := validateUserModuleAccess(userModuleAccessRole.GetUserModuleRole()); err != nil {
		log.Printf("User does not have permission to edit module: %v", err)
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

	// delete the module
	err = c.moduleRepo.DeleteModule(ctx, tx, req.GetModuleId())
	if err != nil {
		log.Printf("Failed to delete module: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("Module deleted successfully")
	return &mpb.DeleteModuleResponse{}, nil
}

func (c *ModuleController) SetUserModuleFavourite(ctx context.Context, req *mpb.SetUserModuleFavouriteRequest) (*mpb.SetUserModuleFavouriteResponse, error) {
	// validate input
	if req.GetUserId() == 0 {
		log.Println("User ID is required")
		return nil, status.Error(codes.InvalidArgument, "User ID is required")
	} else if req.GetModuleId() == 0 {
		log.Println("Module ID is required")
		return nil, status.Error(codes.InvalidArgument, "Module ID is required")
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

	// set user module favourite
	err = c.userModuleMapRepo.PutUserModuleMappingFavourite(ctx, tx, req.GetUserId(), req.GetModuleId(), req.GetIsFavourite())
	if err != nil {
		log.Printf("Failed to set user module favourite: %v", err)
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}

	log.Println("User module favourite set successfully")
	return &mpb.SetUserModuleFavouriteResponse{}, nil
}

func parseAndValidateModuleName(moduleName string) (string, error) {
	// strip leading and trailing whitespace
	moduleName = strings.TrimSpace(moduleName)
	if len(moduleName) < 3 {
		return "", status.Error(codes.InvalidArgument, "Module name must be at least 3 characters")
	} else if len(moduleName) > 100 {
		return "", status.Error(codes.InvalidArgument, "Module name must be less than 100 characters")
	} else{
		return moduleName, nil
	}
}

func validateUserModuleAccess(userModuleAccessRole common.UserModuleRole) error {
	if userModuleAccessRole != common.UserModuleRole_USER_MODULE_ROLE_ADMIN && userModuleAccessRole != common.UserModuleRole_USER_MODULE_ROLE_OWNER && userModuleAccessRole != common.UserModuleRole_USER_MODULE_ROLE_EDITOR {
		return status.Error(codes.PermissionDenied, "User does not have permission to edit module")
	}
	return nil
}
