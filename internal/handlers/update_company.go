package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/akolybelnikov/xm-exercise/internal/models"
)

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var req models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update company
	if err := h.service.UpdateCompany(r.Context(), &req); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusAccepted)
}
