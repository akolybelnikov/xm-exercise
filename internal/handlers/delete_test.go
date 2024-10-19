package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"github.com/akolybelnikov/xm-exercise/internal/handlers"
)

func TestDeleteRequest(t *testing.T) {
	tests := map[string]struct {
		urlParam string
		wantCode int
	}{
		"ValidId":    {urlParam: "/123abc", wantCode: http.StatusAccepted},
		"EmptyId":    {urlParam: "", wantCode: http.StatusNotFound},
		"EmptySlash": {urlParam: "/", wantCode: http.StatusNotFound},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			url := "/delete" + tt.urlParam
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				t.Fatalf("could not create url: %v", err)
			}

			rec := httptest.NewRecorder()
			h := handlers.NewHandler()
			r := chi.NewRouter()
			r.Delete("/delete/{id}", h.Delete)

			r.ServeHTTP(rec, req)

			if status := rec.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantCode)
			}
		})
	}
}
