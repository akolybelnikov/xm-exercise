package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/akolybelnikov/xm-exercise/internal/models"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse request
	var req models.CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validate.Struct(req); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest)+": "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create company
	id, err := h.service.CreateCompany(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return id and code in response
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(id))
	if err != nil {
		return
	}
}
