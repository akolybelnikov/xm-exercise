package app

import (
	"github.com/akolybelnikov/xm-exercise/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	h := handlers.NewHandler()

	r.Route("/api/v1/companies", func(r chi.Router) {
		r.Post("/create", h.Create)
		r.Patch("/update", h.Patch)
		r.Delete("/delete/{id}", h.Delete)
		r.Get("/{id}", h.Get)
	})

	return r
}
