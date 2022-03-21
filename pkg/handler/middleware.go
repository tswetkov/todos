package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
)

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userIdCtx)
	if !ok {
		responseWithError(c, http.StatusInternalServerError, "user id was not found")
		return 0, errors.New("user id was not found")
	}

	id, ok := userId.(int)
	if !ok {
		return 0, errors.New("invalid type of user id")
	}

	return id, nil
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		responseWithError(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	splitHeader := strings.Split(header, " ")
	if len(splitHeader) != 2 {
		responseWithError(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(splitHeader[1])
	if err != nil {
		responseWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtx, userId)
}
