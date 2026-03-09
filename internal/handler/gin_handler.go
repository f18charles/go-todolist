package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"todolist/internal/service"
)

type GinHandler struct {
	Service *service.TodoService
}

func NewGinHandler(s *service.TodoService) *GinHandler {
	return &GinHandler{Service: s}
}

func (h *GinHandler) Create(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}

	c.BindJSON(&req)

	todo, err := h.Service.Create(req.Title)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (h *GinHandler) List(c *gin.Context) {
	todos, err := h.Service.List()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, todos)
}

func (h *GinHandler) Complete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.Complete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *GinHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
