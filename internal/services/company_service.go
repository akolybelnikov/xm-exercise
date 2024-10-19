package services

import (
	"context"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
)

type CompanyService interface {
	GetCompanyByID(ctx context.Context, id string) (*db.Company, error)
	CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error)
	UpdateCompany(ctx context.Context, request *models.UpdateRequest) error
	DeleteCompany(ctx context.Context, id string) error
}

// CompanyDataService is a service that provides methods for managing company data in the domain layer.
type CompanyDataService struct {
	repo repository.CompanyRepository
}

// NewCompanyDataService creates a new CompanyDataService with the provided repository.
func NewCompanyDataService(repo repository.CompanyRepository) *CompanyDataService {
	return &CompanyDataService{repo: repo}
}

// GetCompanyByID returns a company by its ID.
func (s *CompanyDataService) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	return s.repo.GetCompanyByID(ctx, id)
}

// CreateCompany creates a new company.
func (s *CompanyDataService) CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error) {
	return s.repo.CreateCompany(ctx, request)
}

// UpdateCompany updates a company.
func (s *CompanyDataService) UpdateCompany(ctx context.Context, request *models.UpdateRequest) error {
	return s.repo.UpdateCompany(ctx, request)
}

// DeleteCompany deletes a company by its ID.
func (s *CompanyDataService) DeleteCompany(ctx context.Context, id string) error {
	return s.repo.DeleteCompany(ctx, id)
}
