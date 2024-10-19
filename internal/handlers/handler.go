package handlers

import "github.com/go-playground/validator/v10"

type Handler struct {
	validate *validator.Validate
}

func NewHandler() *Handler {
	return &Handler{
		validate: validator.New(),
	}
}
