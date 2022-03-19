package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {}

func (h *Handler) getAllList(c *gin.Context) {
	userId, _ := c.Get(userIdCtx)

	c.JSON(http.StatusOK, map[string]int{
		"userId": userId.(int),
	})
}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
