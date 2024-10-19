package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/handlers"
)

func TestPatchRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCode int
	}{
		{
			name: "valid url",
			input: `{
				"id": "123abc",
				"name": "Test Corp",
				"description": "This is a test corporation",
				"employees": 100,
				"registered": "true",
				"company_type": "Corporations"
			}`,
			wantCode: http.StatusOK,
		},
		{
			name:     "empty url",
			input:    `{}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "invalid company type",
			input: `{
				"id": "456def",
				"name": "Invalid Corp",
				"description": "Corporation with invalid type",
				"employees": 50,
				"registered": "false",
				"company_type": "InvalidType"
			}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "name too long",
			input: `{
				"id": "789ghi",
				"name": "A very long corporation name that exceeds the character limit",
				"description": "This is a test corporation",
				"employees": 100,
				"registered": "true",
				"company_type": "Corporations"
			}`,
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(tt.input))
			if err != nil {
				t.Fatalf("could not create url: %v", err)
			}

			rec := httptest.NewRecorder()
			h := handlers.NewHandler()
			handler := http.HandlerFunc(h.Patch)

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantCode)
			}
		})
	}
}
