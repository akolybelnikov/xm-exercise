package repository_test

import (
	"context"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/stretchr/testify/require"
)

// TestDeleteCompany tests the DeleteCompany method.
func TestDeleteCompany(t *testing.T) {
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

	// Delete the company
	err = repo.DeleteCompany(context.Background(), id)
	require.NoError(t, err)

	// Get the company
	_, err = repo.GetCompanyByID(context.Background(), id)
	require.Error(t, err)
	require.Equal(t, "company not found", err.Error())
}

// TestDeleteCompanyNotFound tests the DeleteCompany method with a non-existent ID.
func TestDeleteCompanyNotFound(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Delete the company with a valid UUID but that does not exist in the database
	companyID := "00000000-0000-0000-0000-000000000000"
	err := repo.DeleteCompany(context.Background(), companyID)
	require.NoError(t, err)
}
