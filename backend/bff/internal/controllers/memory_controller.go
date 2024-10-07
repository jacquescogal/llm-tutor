package controllers

import (
	"bff/internal/proto/memory"
	"bff/internal/services"

	"github.com/gin-gonic/gin"
)

type MemoryController struct{
    memoryService *services.MemoryService
}

func NewMemoryController(memoryService *services.MemoryService) *MemoryController {
	return &MemoryController{memoryService: memoryService}
}

func (c *MemoryController) CreateMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req memory.CreateMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.memoryService.CreateMemory(ctx, &req)
}

func (c *MemoryController) GetMemoryById(ctx *gin.Context) (*memory.GetMemoryByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req memory.GetMemoryByIdRequest
	ctx.Bind(&req)
	req.UserId = userId

	return c.memoryService.GetMemoryById(ctx, &req)
}

func (c *MemoryController) GetMemoriesByDocId(ctx *gin.Context) (*memory.GetMemoriesByDocIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req memory.GetMemoriesByDocIdRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.memoryService.GetMemoriesByDocId(ctx, &req)
}

func (c *MemoryController) GetMemoriesByMemoryTitleSearch(ctx *gin.Context) (*memory.GetMemoriesByMemoryTitleSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	var req memory.GetMemoriesByMemoryTitleSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.memoryService.GetMemoriesByMemoryTitleSearch(ctx, &req)
}

func (c *MemoryController) UpdateMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req memory.UpdateMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.memoryService.UpdateMemory(ctx, &req)
}

func (c *MemoryController) DeleteMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req memory.DeleteMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.memoryService.DeleteMemory(ctx, &req)
}