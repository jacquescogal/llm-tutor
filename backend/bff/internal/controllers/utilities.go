package controllers

import (
	"bff/internal/proto/authenticator"
	"bff/internal/proto/common"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	QUERY_ORDER_BY_FIELD = "sort_by" // created_at, updated_at, title, etc
	QUERY_ORDER_BY_DIRECTION = "order" // asc or desc
	QUERY_PAGE_NUMBER = "page_number"
	QUERY_PAGE_SIZE = "page_size"
)

type QueryItems struct{
	PageNumber uint32
	OrderByField common.ORDER_BY_FIELD
	OrderByDirection common.ORDER_BY_DIRECTION
}

func NewQueryItems(ctx *gin.Context) (*QueryItems, error) {
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
	return &QueryItems{
		PageNumber: pageNumber,
		OrderByField: common.ORDER_BY_FIELD(orderByField),
		OrderByDirection: common.ORDER_BY_DIRECTION(orderByDirection),
	}, nil
}

func getUserSession(ctx *gin.Context) (*authenticator.UserSession, error) {
	userSession, ok := ctx.Get("user_session")
	if !ok {
		return nil, fmt.Errorf("user session not found")
	}
	return userSession.(*authenticator.UserSession), nil
}

func getUint32FromString(value string) (uint32, error) {
	i, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", value)
	}
	return uint32(i), nil
}

func getInt64FromString( value string) (int64, error) {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", value)
	}
	return i, nil
}

func getUint64FromString(value string) (uint64, error) {
	i, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", value)
	}
	return i, nil
}

func getInt32FromString(value string) (int32, error) {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", value)
	}
	return int32(i), nil
}
func getBoolFromString(value string) (bool, error) {
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("invalid %s", value)
	}
	return b, nil
}