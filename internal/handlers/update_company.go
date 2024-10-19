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

	// Validate url
	if err := h.validate.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return response
	w.WriteHeader(http.StatusOK)
}
