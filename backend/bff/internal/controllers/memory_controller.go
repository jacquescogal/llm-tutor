package controllers

import (
	"bff/internal/proto/memory"
	"bff/internal/services"
	"os"

	"github.com/gin-gonic/gin"
)

type MemoryController struct{
    memoryService *services.MemoryService
	pagseSize uint32
}

func NewMemoryController(memoryService *services.MemoryService) *MemoryController {
	pageSizeString := os.Getenv("MEMORY_PAGE_SIZE")
	if pageSizeString == "" {
		// fallback to default value
		pageSizeString = "10"
	}
	pageSize, err := getUint32FromString(pageSizeString)
	if err != nil {
		// fatal error on start up
		panic(err)
	}
	return &MemoryController{memoryService: memoryService, pagseSize: pageSize}
}

func (c *MemoryController) CreateMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	var req memory.CreateMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.DocId = documentId
	return c.memoryService.CreateMemory(ctx, &req)
}

func (c *MemoryController) GetMemoryById(ctx *gin.Context) (*memory.GetMemoryByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	memoryId,err := c.getMemoryIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	var req memory.GetMemoryByIdRequest
	ctx.Bind(&req)
	req.UserId = userId

	req.ModuleId = moduleId
	req.DocId = documentId
	req.MemoryId = memoryId
	return c.memoryService.GetMemoryById(ctx, &req)
}

func (c *MemoryController) GetMemoriesByDocId(ctx *gin.Context) (*memory.GetMemoriesByDocIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	queryItems,err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}

	var req memory.GetMemoriesByDocIdRequest
	ctx.Bind(&req)
	req.UserId = userId

	req.ModuleId = moduleId
	req.DocId = documentId
	req.PageNumber = queryItems.PageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = queryItems.OrderByField
	req.OrderByDirection = queryItems.OrderByDirection
	return c.memoryService.GetMemoriesByDocId(ctx, &req)
}

func (c *MemoryController) GetMemoriesByMemoryTitleSearch(ctx *gin.Context) (*memory.GetMemoriesByMemoryTitleSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return nil, err
	}
	queryItems,err := NewQueryItems(ctx)
	if err != nil {
		return nil, err
	}

	var req memory.GetMemoriesByMemoryTitleSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.DocId = documentId
	req.PageNumber = queryItems.PageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = queryItems.OrderByField
	req.OrderByDirection = queryItems.OrderByDirection
	return c.memoryService.GetMemoriesByMemoryTitleSearch(ctx, &req)
}

func (c *MemoryController) UpdateMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	memoryId,err := c.getMemoryIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	var req memory.UpdateMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.DocId = documentId
	req.MemoryId = memoryId
	return c.memoryService.UpdateMemory(ctx, &req)
}

func (c *MemoryController) DeleteMemory(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	moduleId,err := c.getModuleIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	documentId,err := c.getDocumentIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	memoryId,err := c.getMemoryIdFromContextParams(ctx)
	if err != nil {
		return err
	}
	var req memory.DeleteMemoryRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	req.DocId = documentId
	req.MemoryId = memoryId
	return c.memoryService.DeleteMemory(ctx, &req)
}


func (c *MemoryController) getModuleIdFromContextParams(ctx *gin.Context) (uint64, error) {
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return 0, err
	}
	return moduleId, nil
}

func (c *MemoryController) getDocumentIdFromContextParams(ctx *gin.Context) (uint64, error) {
	documentId, err := getUint64FromString(ctx.Param("document_id"))
	if err != nil {
		return 0, err
	}
	return documentId, nil
}

func (c *MemoryController) getMemoryIdFromContextParams(ctx *gin.Context) (uint64, error) {
	documentId, err := getUint64FromString(ctx.Param("memory_id"))
	if err != nil {
		return 0, err
	}
	return documentId, nil
}