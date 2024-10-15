package controllers

import (
	"bff/internal/proto/common"
	"bff/internal/proto/module"
	"bff/internal/services"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type ModuleController struct{
    moduleService *services.ModuleService
	pagseSize uint32
}

func NewModuleController(moduleService *services.ModuleService) *ModuleController {
	pageSizeString := os.Getenv("MODULE_PAGE_SIZE")
	if pageSizeString == "" {
		// fallback to default value
		pageSizeString = "10"
	}
	pageSize, err := getUint32FromString(pageSizeString)
	if err != nil {
		// fatal error on start up
		panic(err)
	}
	return &ModuleController{moduleService: moduleService, pagseSize: pageSize}
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

func (c *ModuleController) GetPublicModules(ctx *gin.Context) (*module.GetPublicModulesResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	pageNumber, err := getUint32FromString(ctx.DefaultQuery(QUERY_PAGE_NUMBER, "1"))
	if err != nil {
		return nil, err
	}
	orderByField, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_FIELD, "1"))
	if err != nil {
		return nil, err
	}
	orderByDirection, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_DIRECTION, "1"))
	if err != nil {
		return nil, err
	}
	req := module.GetPublicModulesRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.moduleService.GetPublicModules(ctx, &req)
}

func (c *ModuleController) GetPrivateModulesByUserId(ctx *gin.Context) (*module.GetPrivateModulesByUserIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	pageNumber, err := getUint32FromString(ctx.DefaultQuery(QUERY_PAGE_NUMBER, "1"))
	if err != nil {
		return nil, err
	}
	orderByField, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_FIELD, "1"))
	if err != nil {
		return nil, err
	}
	orderByDirection, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_DIRECTION, "1"))
	if err != nil {
		return nil, err
	}
	req := module.GetPrivateModulesByUserIdRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.moduleService.GetPrivateModulesByUserId(ctx, &req)
}

func (c *ModuleController) GetFavouriteModulesByUserId(ctx *gin.Context) (*module.GetFavouriteModulesByUserIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	pageNumber, err := getUint32FromString(ctx.DefaultQuery(QUERY_PAGE_NUMBER, "1"))
	if err != nil {
		return nil, err
	}
	orderByField, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_FIELD, "1"))
	if err != nil {
		return nil, err
	}
	orderByDirection, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_DIRECTION, "1"))
	if err != nil {
		return nil, err
	}
	req := module.GetFavouriteModulesByUserIdRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.moduleService.GetFavouriteModulesByUserId(ctx, &req)
}

func (c *ModuleController) GetModuleById(ctx *gin.Context) (*module.GetModuleByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return nil, err
	}
	req := module.GetModuleByIdRequest{
		UserId: userId,
		ModuleId: moduleId,
	}
	return c.moduleService.GetModuleById(ctx, &req)
}

func (c *ModuleController) GetModulesBySubjectId(ctx *gin.Context) (*module.GetModulesBySubjectIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	subjectId, err := getUint64FromString(ctx.Param("subject_id"))
	if err != nil {
		return nil, err
	}
	fmt.Println("subjectId", subjectId)
	userId := userSession.UserId
	pageNumber, err := getUint32FromString(ctx.DefaultQuery(QUERY_PAGE_NUMBER, "1"))
	if err != nil {
		return nil, err
	}
	orderByField, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_FIELD, "1"))
	if err != nil {
		return nil, err
	}
	orderByDirection, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_DIRECTION, "1"))
	if err != nil {
		return nil, err
	}
	req :=  module.GetModulesBySubjectIdRequest{
		UserId: userId,
		SubjectId: subjectId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.moduleService.GetModulesBySubjectId(ctx, &req)
}

func (c *ModuleController) GetModulesByNameSearch(ctx *gin.Context) (*module.GetModulesByNameSearchResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	pageNumber, err := getUint32FromString(ctx.DefaultQuery(QUERY_PAGE_NUMBER, "1"))
	if err != nil {
		return nil, err
	}
	orderByField, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_FIELD, "1"))
	if err != nil {
		return nil, err
	}
	orderByDirection, err := getInt32FromString(ctx.DefaultQuery(QUERY_ORDER_BY_DIRECTION, "1"))
	if err != nil {
		return nil, err
	}
	var req module.GetModulesByNameSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.PageNumber = pageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = common.ORDER_BY_FIELD(orderByField)
	req.OrderByDirection = common.ORDER_BY_DIRECTION(orderByDirection)

	return c.moduleService.GetModulesByNameSearch(ctx, &req)
}

func (c *ModuleController) UpdateModule(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.UpdateModuleRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	return c.moduleService.UpdateModule(ctx, &req)
}

func (c *ModuleController) DeleteModule(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.DeleteModuleRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	return c.moduleService.DeleteModule(ctx, &req)
}

func (c *ModuleController) SetUserModuleFavourite(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	moduleId, err := getUint64FromString(ctx.Param("module_id"))
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req module.SetUserModuleFavouriteRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.ModuleId = moduleId
	return c.moduleService.SetUserModuleFavourite(ctx, &req)
}