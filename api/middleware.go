package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
)

func checkToken(token string) (bool, error) {
	if token == "" {
		return false, errors.New("No token provided")
	}
	return true, nil
}

func AuthMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	accessToken := fields[1]
	valid, err := checkToken(accessToken)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}
}
