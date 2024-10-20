package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/akolybelnikov/xm-exercise/internal/repository"

	"github.com/golang-jwt/jwt/v4"
)

type LoginHandler struct {
	JwtKey         []byte
	expirationTime int
	repo           repository.CompanyRepository
	validate       *validator.Validate
}

func NewLoginHandler(jwtKey string, exp int, repo repository.CompanyRepository) *LoginHandler {
	return &LoginHandler{
		JwtKey:         []byte(jwtKey),
		expirationTime: exp,
		repo:           repo,
		validate:       validator.New(),
	}
}

// Login handler authenticates the user and issues a JWT token.
func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username" validate:"required,max=15"`
		Password string `json:"password" validate:"required,max=100"`
	}

	// Decode the request body into credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the request
	if err = h.validate.Struct(credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify the user's credentials
	if _, err = h.repo.VerifyUser(r.Context(), credentials.Username, credentials.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Set token expiration time
	expirationTime := time.Now().Add(time.Duration(h.expirationTime) * time.Minute)
	// Create claims with the user's username and expiration time
	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create and sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the token in the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	if err != nil {
		return
	}
}
