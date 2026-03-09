package handler

import (
	"encoding/json"
	"net/http"

	"todolist/internal/service"
)

type StdHandler struct {
	Service *service.TodoService
}

func NewStdHandler(s *service.TodoService) *StdHandler {
	return &StdHandler{Service: s}
}

func (h *StdHandler) Create(w http.ResponseWriter, r *http.Request) {

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

func (h *StdHandler) List(w http.ResponseWriter, r *http.Request) {

	todos, err := h.Service.List()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(todos)
}
