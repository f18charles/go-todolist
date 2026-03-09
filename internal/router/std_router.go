package router

import (
	"net/http"
	"todolist/internal/handler"
)

func NewStdRouter(h *handler.StdHandler) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			h.List(w, r)
			return
		}

		if r.Method == "POST" {
			h.Create(w, r)
			return
		}
	})

	return mux
}
