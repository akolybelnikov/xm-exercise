package app

import (
	"github.com/akolybelnikov/xm-exercise/internal/handlers"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
	"github.com/akolybelnikov/xm-exercise/internal/services"
	"github.com/go-chi/chi/v5"
)

func NewRouter(repo repository.CompanyRepository) *chi.Mux {
	r := chi.NewRouter()
	s := services.NewCompanyDataService(repo)
	h := handlers.NewHandler(s)

	r.Route("/api/v1/companies", func(r chi.Router) {
		r.Post("/create", h.Create)
		r.Patch("/update", h.Patch)
		r.Delete("/delete/{id}", h.Delete)
		r.Get("/{id}", h.Get)
	})

	return r
}
