//go:build integration
// +build integration

package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akolybelnikov/xm-exercise/internal/config"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
)

// TestPostgresCompanyRepository_VerifyUser tests the VerifyUser method.
func TestPostgresCompanyRepository_VerifyUser(t *testing.T) {
	repo, teardown := setUpTestingUser(t)
	defer teardown()

	// Verify a user
	user, err := repo.VerifyUser(context.Background(), "admin", "admin")
	require.NoError(t, err)
	require.NotNilf(t, user, "user should not be nil")
}

func setUpTestingUser(t *testing.T) (repository.CompanyRepository, func()) {
	cfg := &config.DBConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Name:     "postgres",
		SSLMode:  "disable",
	}
	repo, err := repository.NewPostgresCompanyRepository(cfg)
	require.NoError(t, err)

	ctx := context.Background()

	// create the table in the database
	_, ddlErr := repo.Pool.Exec(ctx, `
	DROP TABLE IF EXISTS company CASCADE;
	DROP TABLE IF EXISTS users;
	DROP EXTENSION IF EXISTS "uuid-ossp";
	DROP EXTENSION IF EXISTS pgcrypto;
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
	CREATE EXTENSION IF NOT EXISTS pgcrypto;
	
	CREATE TABLE IF NOT EXISTS company
	(
		id             UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
		name           VARCHAR(15) UNIQUE NOT NULL,
		description    VARCHAR(3000),
		employee_count INT                NOT NULL,
		registered     BOOLEAN            NOT NULL,
		type           VARCHAR(20)        NOT NULL CHECK ( type IN ('Corporations', 'NonProfit', 'Cooperative',
																	'Sole Proprietorship') )
	);
	
	CREATE TABLE IF NOT EXISTS users
	(
		id         SERIAL PRIMARY KEY,
		username   VARCHAR(15) UNIQUE NOT NULL,
		password   VARCHAR(100)       NOT NULL,
		email      VARCHAR(100)       NOT NULL,
		company_id UUID               REFERENCES company (id) ON DELETE SET NULL
	);
	
	INSERT INTO company (name, description, employee_count, registered, type)
	VALUES ('Apple',
			'Apple Inc. is an American multinational technology company that specializes in consumer electronics, computer software, and online services.',
			147000, TRUE, 'Corporations');
	INSERT INTO users (username, password, email, company_id)
	VALUES ('admin', crypt('admin', gen_salt('bf')), 'admin@apple.inc', (SELECT id FROM company WHERE name = 'Apple'));
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
