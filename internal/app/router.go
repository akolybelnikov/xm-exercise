package app

import (
	"github.com/akolybelnikov/xm-exercise/internal/auth"
	"github.com/akolybelnikov/xm-exercise/internal/handlers"
	"github.com/akolybelnikov/xm-exercise/internal/kafka"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
	"github.com/akolybelnikov/xm-exercise/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type RouterConfig struct {
	Producer kafka.MutationProducer
	Repo     repository.CompanyRepository
	Topic    string
	Secret   string
	Exp      int
}

func NewRouter(cfg *RouterConfig) *chi.Mux {
	r := chi.NewRouter()
	// Use the Logger middleware to log requests to the console
	r.Use(middleware.Logger)
	// Use the Recoverer middleware to recover from panics anywhere in the stack
	r.Use(middleware.Recoverer)

	// Register the login handler
	a := auth.NewLoginHandler(cfg.Secret, cfg.Exp, cfg.Repo)
	r.Post("/login", a.Login)

	s := services.NewCompanyDataService(cfg.Producer, cfg.Topic, cfg.Repo)
	h := handlers.NewHandler(s)

	// Register the company handlers
	r.Route("/api/v1/companies", func(r chi.Router) {
		r.Use(a.JWTAuthMiddleware)
		r.Post("/create", h.Create)
		r.Patch("/update", h.Patch)
		r.Delete("/delete/{id}", h.Delete)
		r.Get("/{id}", h.Get)
	})

	return r
}
