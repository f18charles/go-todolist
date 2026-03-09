package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"todolist/internal/service"
)

type ChiHandler struct {
	Service *service.TodoService
}

func NewChiHandler(s *service.TodoService) *ChiHandler {
	return &ChiHandler{Service: s}
}

func (h *ChiHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Title string `json:"title"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	todo, err := h.Service.Create(req.Title)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (h *ChiHandler) List(w http.ResponseWriter, r *http.Request) {

	todos, err := h.Service.List()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
