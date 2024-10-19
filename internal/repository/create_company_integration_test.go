//go:build integration
// +build integration

package repository_test

import (
	"context"
	"testing"

	"github.com/akolybelnikov/xm-exercise/internal/models"

	"github.com/stretchr/testify/require"

	"github.com/akolybelnikov/xm-exercise/internal/config"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
)

func setUpTestingDB(t *testing.T) (repository.CompanyRepository, func()) {
	cfg := &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Name:     "postgres",
		SSLMode:  "disable",
	}
	repo, err := repository.NewPGXCompanyRepository(cfg)
	require.NoError(t, err)

	ctx := context.Background()

	// create the table in the database
	_, ddlErr := repo.Pool.Exec(ctx, `
	DROP TABLE IF EXISTS company;	
	DROP EXTENSION IF EXISTS "uuid-ossp";
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE TABLE IF NOT EXISTS company (
	  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	  name VARCHAR(15) UNIQUE NOT NULL,
	  description VARCHAR(3000),
	  employee_count INT NOT NULL,
	  registered BOOLEAN NOT NULL,
	  type VARCHAR(20) NOT NULL CHECK ( type IN ('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship') )
	);
`)
	require.NoError(t, ddlErr)

	// begin transaction
	tx, err := repo.Pool.Begin(ctx)
	require.NoError(t, err)

	// Override the repository's db with the transaction
	repo.DB = repo.DB.WithTx(tx)

	teardown := func() {
		err2 := tx.Rollback(ctx)
		require.NoError(t, err2)

		repo.Pool.Close()
	}

	return repo, teardown
}

// TestCreateCompany tests the CreateCompany method.
func TestCreateCompany(t *testing.T) {
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

	// Get the company by its ID
	company, err := repo.GetCompanyByID(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, "Test Co.", company.Name)
	require.Equal(t, "This is a test company.", company.Description.String)
	require.Equal(t, int32(100), company.EmployeeCount)
	require.True(t, company.Registered)
	require.Equal(t, "Corporations", company.Type)
}

// TestCreateCompanyInvalidType tests the CreateCompany method with an invalid company type.
func TestCreateCompanyInvalidType(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Create a new company with an invalid company type
	_, err := repo.CreateCompany(context.Background(), &models.CreateRequest{
		Name:          "Test Co.",
		Description:   "This is a test company.",
		EmployeeCount: 100,
		Registered:    "true",
		CompanyType:   "Invalid",
	})
	require.Error(t, err)
}

// TestCreateCompanyNameTooLong tests the CreateCompany method with a name that is too long.
func TestCreateCompanyNameTooLong(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Create a new company with a name that is too long
	_, err := repo.CreateCompany(context.Background(), &models.CreateRequest{
		Name:          "A very long company name that exceeds the character limit",
		Description:   "This is a test company.",
		EmployeeCount: 100,
		Registered:    "true",
		CompanyType:   "Corporations",
	})
	require.Error(t, err)
}

// TestCreateCompanyNonUniqueName tests the CreateCompany method with a non-unique name.
func TestCreateCompanyNonUniqueName(t *testing.T) {
	repo, teardown := setUpTestingDB(t)
	defer teardown()

	// Create a new company
	_, err := repo.CreateCompany(context.Background(), &models.CreateRequest{
		Name:          "Test Co.",
		Description:   "This is a test company.",
		EmployeeCount: 100,
		Registered:    "true",
		CompanyType:   "Corporations",
	})
	require.NoError(t, err)

	// Create a new company with the same name
	_, err = repo.CreateCompany(context.Background(), &models.CreateRequest{
		Name:          "Test Co.",
		Description:   "This is a test company.",
		EmployeeCount: 100,
		Registered:    "true",
		CompanyType:   "Corporations",
	})
	require.Error(t, err)
}
