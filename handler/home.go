package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type homeHandler struct {
}

func NewHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (h *homeHandler) GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
