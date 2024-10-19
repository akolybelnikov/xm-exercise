package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// Parse URL parameters
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	// Return response
	w.WriteHeader(http.StatusOK)
}
