package handlers

import (
	"github.com/akolybelnikov/xm-exercise/internal/services"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	service  services.CompanyService
	validate *validator.Validate
}

func NewHandler(service services.CompanyService) *Handler {
	return &Handler{
		service:  service,
		validate: validator.New(),
	}
}
