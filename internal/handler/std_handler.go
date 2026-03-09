package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-todolist/internal/service"
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

// extractID pulls the last path segment, e.g. /todos/abc-123 → "abc-123"
func extractID(r *http.Request) string {
	parts := strings.Split(r.URL.Path, "/")
	return parts[len(parts)-1]
}

func (h *StdHandler) Complete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)

	if err := h.Service.Complete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *StdHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := extractID(r)

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
