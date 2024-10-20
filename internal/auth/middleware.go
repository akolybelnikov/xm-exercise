package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (h *LoginHandler) JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// The token normally comes in a format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		// Initialize a new instance of Claims to decode the token
		claims := &Claims{}

		// Parse the token and verify it
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(_ *jwt.Token) (interface{}, error) {
			return h.JwtKey, nil
		})
		if err != nil {
			log.Println("Error parsing token:", err)
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Proceed with the request if the token is valid
		next.ServeHTTP(w, r)
	})
}
