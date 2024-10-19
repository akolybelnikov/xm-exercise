package handlers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akolybelnikov/xm-exercise/db"

	"github.com/akolybelnikov/xm-exercise/internal/models"

	"github.com/stretchr/testify/mock"

	"github.com/akolybelnikov/xm-exercise/internal/handlers"
)

type mockCompanyService struct {
	mock.Mock
}

func (m *mockCompanyService) CreateCompany(ctx context.Context, req *models.CreateRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

func (m *mockCompanyService) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*db.Company), args.Error(1)
}

func (m *mockCompanyService) UpdateCompany(ctx context.Context, req *models.UpdateRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockCompanyService) DeleteCompany(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

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
			s := new(mockCompanyService)
			h := handlers.NewHandler(s)
			handler := http.HandlerFunc(h.Create)

			handler.ServeHTTP(rec, req)

			if status := rec.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantCode)
			}
		})
	}
}
