package handlers

import (
	"encoding/json"
	"net/http"
)

type PatchRequest struct {
	ID           string `json:"id" validate:"required"`
	Name         string `json:"name" validate:"max=15"`
	Description  string `json:"description" validate:"max=3000"`
	NumEmployees int    `json:"employees"`
	Registered   bool   `json:"registered"`
	CompanyType  string `json:"company_type" validate:"oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship'"` //nolint:lll
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	// Parse url
	var req PatchRequest
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
