package repository

import (
	"context"

	"github.com/akolybelnikov/xm-exercise/db"
)

func (r *PostgresCompanyRepository) VerifyUser(ctx context.Context, username, password string) (*db.User, error) {
	user, err := r.DB.VerifyUser(ctx, db.VerifyUserParams{
		Username: username,
		Crypt:    password,
	})
	if err != nil {
		return nil, err
	}

	return &user, err
}
