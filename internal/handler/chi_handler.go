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

func (h *ChiHandler) Complete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.Service.Complete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChiHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
