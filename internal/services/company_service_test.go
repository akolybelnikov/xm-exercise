package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/services"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCompanyRepo is a mock implementation of the CompanyRepository interface for testing purposes.
type MockCompanyRepo struct {
	mock.Mock
}

// GetCompanyByID retrieves a company by its ID from the repository using context and ID as parameters.
func (m *MockCompanyRepo) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*db.Company), args.Error(1)
}

// CreateCompany creates a new company record in the repository.
// ctx: Context for managing request deadlines, cancellation signals, and other request-scoped values.
// req: CreateRequest containing data necessary to create a new company.
// Returns the ID of the newly created company and an error if something goes wrong.
func (m *MockCompanyRepo) CreateCompany(ctx context.Context, req *models.CreateRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

// UpdateCompany updates the company details in the repository based on the provided UpdateRequest.
// It returns an error if the update operation fails.
func (m *MockCompanyRepo) UpdateCompany(ctx context.Context, req *models.UpdateRequest) error {
	return m.Called(ctx, req).Error(0)
}

// DeleteCompany deletes a company record from the repository by its ID. Returns error if the deletion process fails.
func (m *MockCompanyRepo) DeleteCompany(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}

// TestCompanyByID verifies the functionality of retrieving company details by ID using different test cases.
func TestCompanyByID(t *testing.T) {
	// Create a new UUID.
	uuidID, err := uuid.NewUUID()
	require.NoError(t, err)
	// Convert the UUID to a pgtype.UUID.
	var pgUUID pgtype.UUID
	err = pgUUID.Scan(uuidID.String())
	require.NoError(t, err)
	// Convert string to pgtype.Text.
	var pgText pgtype.Text
	err = pgText.Scan("Test Company")
	require.NoError(t, err)

	testCases := []struct {
		name   string
		id     string
		result *db.Company
		err    error
	}{
		{
			name: "exists",
			id:   uuidID.String(),
			result: &db.Company{
				ID:            pgUUID,
				Name:          "TestCompany",
				Description:   pgText,
				EmployeeCount: int32(100),
				Registered:    true,
				Type:          "Corporations",
			},
			err: nil,
		},
		{
			name:   "not exists",
			id:     uuidID.String(),
			result: nil,
			err:    errors.New("company not found")},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// MockCompanyRepo is a mock implementation of the CompanyRepository interface.
			m := new(MockCompanyRepo)
			// NewCompanyService creates a new CompanyService.
			s := services.NewCompanyService(m)
			m.On("GetCompanyByID", mock.Anything, test.id).Return(test.result, test.err)
			resp, err2 := s.GetCompanyByID(context.Background(), test.id)
			assert.Equal(t, test.result, resp)
			assert.Equal(t, test.err, err2)
			m.AssertExpectations(t)
		})
	}
}

// TestCreateCompany tests the creation of a company using different test cases.
func TestCreateCompany(t *testing.T) {
	testCases := []struct {
		name string
		req  *models.CreateRequest
		id   string
		err  error
	}{
		{
			name: "valid",
			req: &models.CreateRequest{
				Name:          "TestCompany",
				Description:   "Desc",
				EmployeeCount: int32(100),
				Registered:    "false",
				CompanyType:   "Corporations",
			},
			id:  "123",
			err: nil,
		},
		{
			name: "invalid",
			req:  nil,
			id:   "",
			err:  errors.New("invalid request"),
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// MockCompanyRepo is a mock implementation of the CompanyRepository interface.
			m := new(MockCompanyRepo)
			// NewCompanyService creates a new CompanyService.
			s := services.NewCompanyService(m)
			m.On("CreateCompany", mock.Anything, test.req).Return(test.id, test.err)
			resp, err := s.CreateCompany(context.Background(), test.req)
			assert.Equal(t, test.id, resp)
			assert.Equal(t, test.err, err)
			m.AssertExpectations(t)
		})
	}
}

// TestUpdateCompany tests the update functionality of CompanyService.
func TestUpdateCompany(t *testing.T) {
	testCases := []struct {
		name string
		req  *models.UpdateRequest
		err  error
	}{
		{
			name: "valid",
			req: &models.UpdateRequest{
				ID:          "123",
				Name:        "TestCompany",
				Description: "Desc",
			},
			err: nil,
		},
		{"invalid", nil, errors.New("invalid request")},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// MockCompanyRepo is a mock implementation of the CompanyRepository interface.
			m := new(MockCompanyRepo)
			// NewCompanyService creates a new CompanyService.
			s := services.NewCompanyService(m)
			m.On("UpdateCompany", mock.Anything, test.req).Return(test.err)
			err := s.UpdateCompany(context.Background(), test.req)
			assert.Equal(t, test.err, err)
			m.AssertExpectations(t)
		})
	}
}

// TestDeleteCompany tests the DeleteCompany method of the CompanyService.
func TestDeleteCompany(t *testing.T) {
	testCases := []struct {
		name string
		id   string
		err  error
	}{
		{"exists", "123", nil},
		{"not exists", "456", errors.New("company not found")},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// MockCompanyRepo is a mock implementation of the CompanyRepository interface.
			m := new(MockCompanyRepo)
			// NewCompanyService creates a new CompanyService.
			s := services.NewCompanyService(m)
			m.On("DeleteCompany", mock.Anything, test.id).Return(test.err)
			err := s.DeleteCompany(context.Background(), test.id)
			assert.Equal(t, test.err, err)
			m.AssertExpectations(t)
		})
	}
}
