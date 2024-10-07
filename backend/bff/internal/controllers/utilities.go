package controllers

import (
	"bff/internal/proto/authenticator"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

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