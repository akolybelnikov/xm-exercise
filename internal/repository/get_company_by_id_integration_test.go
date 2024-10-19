//go:build integration
// +build integration

package repository_test

import (
	"testing"

	"golang.org/x/net/context"

	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
)

// TestGetCompanyByID tests the GetCompanyByID method.
func TestGetCompanyByID(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Create a new company
	id, err := repo.CreateCompany(context.Background(), &models.CreateRequest{
		Name:          "Test Co.",
		Description:   "This is a test company.",
		EmployeeCount: 100,
		Registered:    "true",
		CompanyType:   "Corporations",
	})
	require.NoError(t, err)

	// Get the company
	company, err := repo.GetCompanyByID(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, "Test Co.", company.Name)
	require.Equal(t, "This is a test company.", company.Description.String)
	require.Equal(t, int32(100), company.EmployeeCount)
	require.True(t, company.Registered)
	require.Equal(t, "Corporations", company.Type)
}

// TestGetCompanyByIDNotFound tests the GetCompanyByID method with a non-existent ID.
func TestGetCompanyByIDNotFound(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Get the company with a valid UUID but that does not exist in the database
	companyID := uuid.New().String()
	_, err := repo.GetCompanyByID(context.Background(), companyID)
	require.Error(t, err)
	require.Equal(t, "company not found", err.Error())
}
