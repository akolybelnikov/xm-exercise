package repository

import (
	"context"
	"errors"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/akolybelnikov/xm-exercise/internal/models"
)

func (r *PostgresCompanyRepository) UpdateCompany(ctx context.Context, request *models.UpdateRequest) error {
	pgUUID, err := ConvertStringToUUID(request.ID)
	if err != nil {
		return errors.New("Error scanning the ID: " + err.Error())
	}

	company, err := r.DB.GetCompanyByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.New("company not found")
		}
		return errors.New("error getting the company: " + err.Error())
	}

	if request.Name != "" {
		company.Name = request.Name
	}

	var pgText pgtype.Text
	description := request.Description
	if description != "" {
		pgText, err = ConvertToText(description)
		if err != nil {
			return errors.New("Error converting the company description: " + err.Error())
		}
		company.Description = pgText
	}

	if request.EmployeeCount != 0 {
		company.EmployeeCount = request.EmployeeCount
	} else if request.EmployeeCount == 0 && company.EmployeeCount != 0 {
		company.EmployeeCount = 0
	}

	if request.Registered == "false" {
		company.Registered = false
	} else if request.Registered == "true" {
		company.Registered = true
	}

	if request.CompanyType != "" {
		company.Type = request.CompanyType
	}

	updateParams := db.UpdateCompanyParams{
		ID:            pgUUID,
		Name:          company.Name,
		Description:   company.Description,
		EmployeeCount: company.EmployeeCount,
		Registered:    company.Registered,
		Type:          company.Type,
	}

	err = r.DB.UpdateCompany(ctx, updateParams)
	if err != nil {
		return errors.New("Error updating the company: " + err.Error())
	}

	return nil
}
