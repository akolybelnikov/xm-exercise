package repository_test

import (
	"context"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/stretchr/testify/require"
)

// TestUpdateCompany tests the UpdateCompany method.
func TestUpdateCompany(t *testing.T) {
	t.Parallel()
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

	// Update the company partially
	err = repo.UpdateCompany(context.Background(), &models.UpdateRequest{
		ID:            id,
		Name:          "New Test Co.",
		Description:   "This is a new test company.",
		EmployeeCount: 200,
	})
	require.NoError(t, err)

	// Get the company
	company, err := repo.GetCompanyByID(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, "New Test Co.", company.Name)
	require.Equal(t, "This is a new test company.", company.Description.String)
	require.Equal(t, int32(200), company.EmployeeCount)
	require.True(t, company.Registered)
	require.Equal(t, "Corporations", company.Type)
}

// TestUpdateCompanyNotFound tests the UpdateCompany method with a non-existent ID.
func TestUpdateCompanyNotFound(t *testing.T) {
	t.Parallel()
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Update the company with a valid UUID but that does not exist in the database
	err := repo.UpdateCompany(context.Background(), &models.UpdateRequest{
		ID: "123e4567-e89b-12d3-a456-426614174000",
	})
	require.Error(t, err)
	require.Equal(t, "company not found", err.Error())
}

// TestUpdateCompanySetCountToZero tests the UpdateCompany method with setting the employee count to zero.
func TestUpdateCompanySetCountToZero(t *testing.T) {
	t.Parallel()
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

	// Update the company with setting the employee count to zero
	err = repo.UpdateCompany(context.Background(), &models.UpdateRequest{
		ID:            id,
		EmployeeCount: 0,
	})
	require.NoError(t, err)

	// Get the company
	company, err := repo.GetCompanyByID(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, int32(0), company.EmployeeCount)
}
