package controllers

import (
	"bff/internal/proto/common"
	"bff/internal/proto/subject"
	"bff/internal/services"
	"os"

	"github.com/gin-gonic/gin"
)

type SubjectController struct{
	subjectService *services.SubjectService
	pagseSize uint32
}

func NewSubjectController(subjectService *services.SubjectService) *SubjectController {
	pageSizeString := os.Getenv("SUBJECT_PAGE_SIZE")
	if pageSizeString == "" {
		// fallback to default value
		pageSizeString = "10"
	}
	pageSize, err := getUint32FromString(pageSizeString)
	if err != nil {
		// fatal error on start up
		panic(err)
	}
	return &SubjectController{subjectService: subjectService, pagseSize: pageSize}
}

func (c *SubjectController) CreateSubject(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	var req subject.CreateSubjectRequest
	ctx.Bind(&req)
	req.UserId = userId
	return c.subjectService.CreateSubject(ctx, &req)
}

func (c *SubjectController) GetPublicSubjects(ctx *gin.Context) (*subject.GetPublicSubjectsResponse, error) {
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
	req := subject.GetPublicSubjectsRequest{
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.subjectService.GetPublicSubjects(ctx, &req)
}

func (c *SubjectController) GetPrivateSubjectsByUserID(ctx *gin.Context) (*subject.GetPrivateSubjectsByUserIdResponse, error) {
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
	req := subject.GetPrivateSubjectsByUserIdRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.subjectService.GetPrivateSubjectsByUserID(ctx, &req)
}

func (c *SubjectController) GetFavouriteSubjectsByUserID(ctx *gin.Context) (*subject.GetFavouriteSubjectsByUserIdResponse, error) {
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
	req := subject.GetFavouriteSubjectsByUserIdRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.subjectService.GetFavouriteSubjectsByUserID(ctx, &req)
}

func (c *SubjectController) GetSubjectByID(ctx *gin.Context) (*subject.GetSubjectByIdResponse, error) {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return nil, err
	}
	userId := userSession.UserId
	subjectId, err := getUint64FromString(ctx.Param("subject_id"))
	if err != nil {
		return nil, err
	}
	var req subject.GetSubjectByIdRequest
	req.UserId = userId
	req.SubjectId = subjectId
	return c.subjectService.GetSubjectByID(ctx, &req)
}

func (c *SubjectController) GetSubjectsByUserID(ctx *gin.Context) (*subject.GetSubjectsByUserIdResponse, error) {
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
	req := subject.GetSubjectsByUserIdRequest{
		UserId: userId,
		PageNumber: pageNumber,
		PageSize: c.pagseSize,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}
	return c.subjectService.GetSubjectsByUserID(ctx, &req)
}

func (c *SubjectController) GetSubjectsByNameSearch(ctx *gin.Context) (*subject.GetSubjectsByNameSearchResponse, error) {
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
	var req subject.GetSubjectsByNameSearchRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.PageNumber = pageNumber
	req.PageSize = c.pagseSize
	req.OrderByField = common.ORDER_BY_FIELD(orderByField)
	req.OrderByDirection = common.ORDER_BY_DIRECTION(orderByDirection)
	return c.subjectService.GetSubjectsByNameSearch(ctx, &req)
}

func (c *SubjectController) UpdateSubject(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	subjectId, err := getUint64FromString(ctx.Param("subject_id"))
	if err != nil {
		return err
	}
	var req subject.UpdateSubjectRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.SubjectId = subjectId
	return c.subjectService.UpdateSubject(ctx, &req)
}

func (c *SubjectController) DeleteSubject(ctx *gin.Context) error {
	userSession,err := getUserSession(ctx)
	if err != nil {
		return err
	}
	userId := userSession.UserId
	subjectId, err := getUint64FromString(ctx.Param("subject_id"))
	if err != nil {
		return err
	}
	var req subject.DeleteSubjectRequest
	ctx.Bind(&req)
	req.UserId = userId
	req.SubjectId = subjectId
	return c.subjectService.DeleteSubject(ctx, &req)
}