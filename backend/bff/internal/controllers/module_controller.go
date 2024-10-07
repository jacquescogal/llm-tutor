package controllers

import (
	"bff/internal/proto/module"
	"bff/internal/services"

	"github.com/gin-gonic/gin"
)

type ModuleController struct{
    moduleService *services.ModuleService
}

func NewModuleController(moduleService *services.ModuleService) *ModuleController {
	return &ModuleController{moduleService: moduleService}
}

func (c *ModuleController) CreateModule(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.CreateModuleRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.moduleService.CreateModule(ctx, &req)
}

func (c *ModuleController) GetModuleById(ctx *gin.Context) (*module.GetModuleByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req module.GetModuleByIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.moduleService.GetModuleById(ctx, &req)
}

func (c *ModuleController) GetModulesBySubjectId(ctx *gin.Context) (*module.GetModulesBySubjectIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req module.GetModulesBySubjectIdRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.moduleService.GetModulesBySubjectId(ctx, &req)
}

func (c *ModuleController) GetModulesByNameSearch(ctx *gin.Context) (*module.GetModulesByNameSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req module.GetModulesByNameSearchRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.moduleService.GetModulesByNameSearch(ctx, &req)
}

func (c *ModuleController) UpdateModule(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.UpdateModuleRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.moduleService.UpdateModule(ctx, &req)
}

func (c *ModuleController) DeleteModule(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.DeleteModuleRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.moduleService.DeleteModule(ctx, &req)
}