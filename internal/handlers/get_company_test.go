package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/stretchr/testify/mock"

	"github.com/go-chi/chi/v5"

	"github.com/akolybelnikov/xm-exercise/internal/handlers"
)

func TestGetRequest(t *testing.T) {
	tests := map[string]struct {
		url      string
		wantCode int
	}{
		"ValidId":    {url: "/123abc", wantCode: http.StatusOK},
		"EmptyId":    {url: "", wantCode: http.StatusNotFound},
		"EmptySlash": {url: "/", wantCode: http.StatusNotFound},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			if err != nil {
				t.Fatalf("could not create url: %v", err)
			}

			rec := httptest.NewRecorder()
			s := new(mockCompanyService)
			h := handlers.NewHandler(s)
			r := chi.NewRouter()
			r.Get("/{id}", h.Get)

			s.On("GetCompanyByID", mock.Anything, "123abc").Return(&db.Company{
				ID:            pgtype.UUID{Bytes: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}},
				Name:          "Test",
				Description:   pgtype.Text{String: "Test description"},
				EmployeeCount: 100,
				Registered:    true,
				Type:          "Corporations",
			}, nil)

			r.ServeHTTP(rec, req)

			if status := rec.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantCode)
			}
		})
	}
}
