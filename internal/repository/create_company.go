package repository

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/akolybelnikov/xm-exercise/db"
)

// CreateCompany creates a new company.
func (r *PGXCompanyRepository) CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error) {
	// convert request.Description to pgtype.Text
	pgText, convertErr := ConvertToText(request.Description)
	if convertErr != nil {
		return "", errors.New("Error converting the company description: " + convertErr.Error())
	}

	registered := false
	if request.Registered == "true" {
		registered = true
	}

	// convert request to CreateCompanyParams
	params := db.CreateCompanyParams{
		Name:          request.Name,
		Description:   pgText,
		EmployeeCount: request.EmployeeCount,
		Registered:    registered,
		Type:          request.CompanyType,
	}

	// attempt to create the company in the database
	uuid, err := r.DB.CreateCompany(ctx, params)
	if err != nil {
		return "", errors.New("Error creating the request: " + err.Error())
	}

	id, convertErr := ConvertUUIDToString(uuid)
	if convertErr != nil {
		return "", errors.New("Error converting the company ID: " + convertErr.Error())
	}

	return id, nil
}

// ConvertToText converts a string to a pgtype.Text.
func ConvertToText(s string) (pgtype.Text, error) {
	var pgText pgtype.Text
	scanErr := pgText.Scan(s)
	if scanErr != nil {
		return pgText, scanErr
	}

	return pgText, nil
}

// ConvertUUIDToString converts a pgtype.UUID to a string.
func ConvertUUIDToString(id pgtype.UUID) (string, error) {
	value, err := id.UUIDValue()
	if err != nil {
		return "", errors.New("Error getting the UUID value: " + err.Error())
	}
	// convert pgtype.UUID to string
	hexString := hex.EncodeToString(value.Bytes[:])

	return hexString, nil
}
