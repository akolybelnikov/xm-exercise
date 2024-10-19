package repository

import (
	"context"
	"errors"
)

func (r *PostgresCompanyRepository) DeleteCompany(ctx context.Context, id string) error {
	pgUUID, convertErr := ConvertStringToUUID(id)
	if convertErr != nil {
		return errors.New("could not convert string to uuid")
	}

	err := r.DB.DeleteCompany(ctx, pgUUID)
	if err != nil {
		return errors.New("could not delete company: " + err.Error())
	}

	return nil
}
