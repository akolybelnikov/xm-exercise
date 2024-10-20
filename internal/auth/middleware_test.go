package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/auth"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuthMiddleware(t *testing.T) {
	jwtKey := []byte("my_secret_key")
	handler := &auth.LoginHandler{JwtKey: jwtKey}

	// Generate a valid token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{Username: "testuser"})
	tokenString, _ := token.SignedString(jwtKey)

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{"NoAuthHeader", "", http.StatusUnauthorized},
		{"InvalidAuthHeader", "Invalid header", http.StatusUnauthorized},
		{"InvalidAuthFormat", "BearerInvalidToken", http.StatusUnauthorized},
		{"InvalidToken", "Bearer invalidToken", http.StatusBadRequest},
		{"ValidToken", "Bearer " + tokenString, http.StatusOK},
	}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("Authorization", tt.authHeader)

			rr := httptest.NewRecorder()
			handler.JWTAuthMiddleware(nextHandler).ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}
