package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/auth"
	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

// Mock repository for testing.
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetCompanyByID(_ context.Context, _ string) (*db.Company, error) {
	// TODO implement me
	panic("implement me")
}

func (m *MockRepository) CreateCompany(_ context.Context, _ *models.CreateRequest) (string, error) {
	// TODO implement me
	panic("implement me")
}

func (m *MockRepository) UpdateCompany(_ context.Context, _ *models.UpdateRequest) error {
	// TODO implement me
	panic("implement me")
}

func (m *MockRepository) DeleteCompany(_ context.Context, _ string) error {
	// TODO implement me
	panic("implement me")
}

func (m *MockRepository) VerifyUser(ctx context.Context, username, password string) (*db.User, error) {
	args := m.Called(ctx, username, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*db.User), args.Error(1)
}

func TestLoginHandler_Login(t *testing.T) {
	tests := []struct {
		name               string
		username           string
		password           string
		expectedStatusCode int
		err                error
	}{
		{
			name:               "valid credentials",
			username:           "validuser",
			password:           "validpass",
			expectedStatusCode: http.StatusOK,
			err:                nil,
		},
		{
			name:               "invalid credentials",
			username:           "invaliduser",
			password:           "invalidpass",
			expectedStatusCode: http.StatusUnauthorized,
			err:                errors.New("invalid credentials"),
		},
		{
			name:               "invalid JSON",
			username:           "",
			password:           "",
			expectedStatusCode: http.StatusBadRequest,
			err:                errors.New("invalid JSON"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request body
			body := map[string]string{
				"username": tt.username,
				"password": tt.password,
			}
			jsonBody, _ := json.Marshal(body)

			req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			mockRepo := new(MockRepository)
			handler := auth.NewLoginHandler("secret", 15, mockRepo)
			h := http.HandlerFunc(handler.Login)

			// Mock the VerifyUser method
			if tt.err != nil {
				mockRepo.On("VerifyUser", mock.Anything, tt.username, tt.password).Return(nil, tt.err)
			} else {
				mockRepo.On("VerifyUser", mock.Anything, tt.username, tt.password).Return(&db.User{}, nil)
			}

			h.ServeHTTP(rr, req)

			require.Equal(t, tt.expectedStatusCode, rr.Code)

			if tt.expectedStatusCode == http.StatusOK {
				var res map[string]string
				err = json.NewDecoder(rr.Body).Decode(&res)
				require.NoError(t, err)
				require.NotEmpty(t, res["token"])

				// Verify token
				token, err2 := jwt.Parse(res["token"], func(_ *jwt.Token) (interface{}, error) {
					return []byte("secret"), nil
				})
				require.NoError(t, err2)
				require.NotNil(t, token)
			}
		})
	}
}
