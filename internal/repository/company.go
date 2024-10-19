package repository

import (
	"context"
	"errors"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository interface {
	GetCompanyByID(id int) (db.Company, error)
	CreateCompany(company db.Company) (int32, error)
	UpdateCompany(company db.Company) error
	DeleteCompany(id int32) error
}

// PGXCompanyRepository is a Postgresql implementation of CompanyRepository using pgx.
type PGXCompanyRepository struct {
	db *db.Queries
}

// NewPGXCompanyRepository creates a new instance of PGXCompanyRepository.
func NewPGXCompanyRepository(cfg *config.DBConfig) (*PGXCompanyRepository, error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.GetDSN())
	if err != nil {
		return nil, errors.New("Error parsing the database config: " + err.Error())
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, errors.New("Error connecting to the database: " + err.Error())
	}
	defer pool.Close()

	queries := db.New(pool)

	return &PGXCompanyRepository{db: queries}, nil
}
