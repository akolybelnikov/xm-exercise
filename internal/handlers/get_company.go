package handlers

import (
	"net/http"

	"github.com/akolybelnikov/xm-exercise/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Parse URL parameters
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	// Get company
	res, err := h.service.GetCompanyByID(r.Context(), id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	company, err := models.FromDBCompany(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusOK)
	_ = render.Render(w, r, company)
}
