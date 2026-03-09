package router

import (
	"github.com/gin-gonic/gin"
	"todolist/internal/handler"
)

func NewGinRouter(h *handler.GinHandler) *gin.Engine {

	r := gin.Default()

	r.GET("/todos", h.List)
	r.POST("/todos", h.Create)

	return r
}
