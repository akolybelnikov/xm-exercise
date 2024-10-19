package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/handlers"
)

func TestCreateRequest(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantCode int
	}{
		{
			name: "CompleteValidRequest",
			input: `{
				"name":"Test Co.",
				"description":"This is a test company.",
				"employee_count":100,
				"registered":"true",
				"company_type":"Corporations"
			}`,
			wantCode: http.StatusCreated,
		},
		{
			name: "CompleteValidRequestNoDescription",
			input: `{
				"name":"Test Co.",
				"employee_count":100,
				"registered":"true",
				"company_type":"Corporations"
			}`,
			wantCode: http.StatusCreated,
		},
		{
			name: "InvalidCompanyType",
			input: `{
				"name":"Test Co.",
				"description":"This is a test company.",
				"employee_count":100,
				"registered":"true",
				"company_type":"Invalid"
			}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "NameTooLong",
			input: `{
				"name":"A very long company name that exceeds the character limit",
				"description": "This is a test company.",
				"employee_count":100,
				"registered":"true",
				"company_type":"Corporations"
			}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "NoNumEmployee_count",
			input: `{
				"name":"Test Co.",
				"registered":"false",
				"company_type":"Corporations"
			}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "WrongNumEmployee_countType",
			input: `{
				"name":"Test Co.",
				"description":"This is a test company.",
				"employee_count":"100",
				"registered":"true",
				"company_type":"Corporations"
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
			handler := http.HandlerFunc(h.Create)

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantCode)
			}
		})
	}
}
