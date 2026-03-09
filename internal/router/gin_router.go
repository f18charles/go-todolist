package router

import (
	"github.com/gin-gonic/gin"
	"go-todolist/internal/handler"
)

func NewGinRouter(h *handler.GinHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", h.List)
	r.POST("/todos", h.Create)
	r.PATCH("/todos/:id", h.Complete)
	r.DELETE("/todos/:id", h.Delete)

	return r
}
