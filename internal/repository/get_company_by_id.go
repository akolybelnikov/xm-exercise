package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/jackc/pgx/v5/pgtype"
)

// GetCompanyByID returns a company by its ID.
func (r *PostgresCompanyRepository) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	// convert string id to pgtype.UUID
	pgUUID, err := ConvertStringToUUID(id)
	if err != nil {
		return nil, errors.New("Error scanning the ID: " + err.Error())
	}

	company, err := r.DB.GetCompanyByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("company not found")
		}
		return nil, errors.New("error getting the company: " + err.Error())
	}

	return &company, nil
}

// ConvertStringToUUID converts a string to a pgtype.UUID.
func ConvertStringToUUID(id string) (pgtype.UUID, error) {
	var pgUUID pgtype.UUID
	err := pgUUID.Scan(id)
	if err != nil {
		return pgUUID, errors.New("Error scanning the ID: " + err.Error())
	}

	return pgUUID, nil
}
