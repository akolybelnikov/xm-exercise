package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateRequest struct {
	Name          string `json:"name" validate:"required,max=15"`
	Description   string `json:"description" validate:"max=3000"`
	EmployeeCount int    `json:"employee_count" validate:"required"`
	Registered    bool   `json:"registered" validate:"required"`
	CompanyType   string `json:"company_type" validate:"required,oneof='Corporations' 'NonProfit' 'Cooperative' 'Sole Proprietorship'"` //nolint:lll
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse url
	var req CreateRequest
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
	w.WriteHeader(http.StatusCreated)
}
