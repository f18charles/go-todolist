package router

import (
	"github.com/go-chi/chi/v5"
	"todolist/internal/handler"
)

func NewChiRouter(h *handler.ChiHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/todos", h.List)
	r.Post("/todos", h.Create)
	r.Patch("/todos/{id}", h.Complete)
	r.Delete("/todos/{id}", h.Delete)

	return r
}
