package services

import (
	"context"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

// GetCompanyByID returns a company by its ID.
func (s *CompanyService) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	return s.repo.GetCompanyByID(ctx, id)
}

// CreateCompany creates a new company.
func (s *CompanyService) CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error) {
	return s.repo.CreateCompany(ctx, request)
}

// UpdateCompany updates a company.
func (s *CompanyService) UpdateCompany(ctx context.Context, request *models.UpdateRequest) error {
	return s.repo.UpdateCompany(ctx, request)
}

// DeleteCompany deletes a company by its ID.
func (s *CompanyService) DeleteCompany(ctx context.Context, id string) error {
	return s.repo.DeleteCompany(ctx, id)
}
