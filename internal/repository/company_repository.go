package repository

import (
	"context"
	"errors"

	"github.com/akolybelnikov/xm-exercise/internal/models"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository interface {
	GetCompanyByID(ctx context.Context, id string) (*db.Company, error)
	CreateCompany(ctx context.Context, req *models.CreateRequest) (string, error)
	UpdateCompany(ctx context.Context, req *models.UpdateRequest) error
	DeleteCompany(ctx context.Context, id string) error
	VerifyUser(ctx context.Context, username, password string) (*db.User, error)
}

// PostgresCompanyRepository is a Postgresql implementation of CompanyRepository using pgx.
type PostgresCompanyRepository struct {
	DB   *db.Queries
	Pool *pgxpool.Pool
}

// NewPostgresCompanyRepository creates a new instance of PostgresCompanyRepository.
func NewPostgresCompanyRepository(cfg *config.DBConfig) (*PostgresCompanyRepository, error) {
	dbConfig, err := pgxpool.ParseConfig(cfg.GetDSN())
	if err != nil {
		return nil, errors.New("Error parsing the database config: " + err.Error())
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, errors.New("Error connecting to the database: " + err.Error())
	}

	queries := db.New(pool)

	return &PostgresCompanyRepository{DB: queries, Pool: pool}, nil
}
