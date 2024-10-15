package controllers

import (
	"bff/internal/proto/generation"
	"bff/internal/services"
	"errors"

	"github.com/gin-gonic/gin"
)

type GenerationController struct{
    generationService *services.GenerationService
}

func NewGenerationController(generationService *services.GenerationService) *GenerationController {
	return &GenerationController{generationService: generationService}
}

func (c *GenerationController) CreateGeneration(ctx *gin.Context) (*generation.CreateGenerationResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil,err
	}
	userId := userSession.UserId
	var req generation.CreateGenerationRequest
	ctx.Bind(&req)
	req.UserId = userId
	if req.ChatHistory == nil {
		return nil,errors.New("ChatHistory is required")
	}else if req.Id == 0 {
		return nil,errors.New("Id is required")
	}else if req.GetIdType() == 0 {
		return nil,errors.New("IdType is required")
	}


	res,err := c.generationService.CreateGeneration(ctx, &req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
// func (c *ModuleController) CreateModule(ctx *gin.Context) error {
// 	userSession,err := getUserSession(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	userId := userSession.UserId
// 	var req module.CreateModuleRequest
// 	ctx.Bind(&req)
// 	req.UserId = userId

// 	return c.moduleService.CreateModule(ctx, &req)
// }
