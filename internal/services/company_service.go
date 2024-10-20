package services

import (
	"context"

	"github.com/akolybelnikov/xm-exercise/internal/kafka"

	"github.com/akolybelnikov/xm-exercise/db"
	"github.com/akolybelnikov/xm-exercise/internal/models"
	"github.com/akolybelnikov/xm-exercise/internal/repository"
)

// CompanyService defines methods for managing companies in the system.
// GetCompanyByID retrieves a company by its ID.
// CreateCompany creates a new company using the provided request data.
// UpdateCompany updates an existing company using the provided request data.
// DeleteCompany removes a company identified by its ID.
type CompanyService interface {
	GetCompanyByID(ctx context.Context, id string) (*db.Company, error)
	CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error)
	UpdateCompany(ctx context.Context, request *models.UpdateRequest) error
	DeleteCompany(ctx context.Context, id string) error
}

// CompanyDataService provides operations related to company management including creation, retrieval, updating,
// and deletion.
type CompanyDataService struct {
	producer kafka.MutationProducer
	repo     repository.CompanyRepository
	topic    string
}

// NewCompanyDataService initializes and returns a new instance of CompanyDataService with the provided producer,
// topic, and repository.
func NewCompanyDataService(
	producer kafka.MutationProducer, topic string, repo repository.CompanyRepository) *CompanyDataService {
	return &CompanyDataService{producer: producer, repo: repo, topic: topic}
}

// GetCompanyByID retrieves a company by its ID.
// ctx: Context for managing request deadlines and other request-scoped values.
// id: The unique identifier string of the company to be retrieved.
// Returns a pointer to db.Company and an error if any issue occurs.
func (s *CompanyDataService) GetCompanyByID(ctx context.Context, id string) (*db.Company, error) {
	return s.repo.GetCompanyByID(ctx, id)
}

// CreateCompany creates a new company based on the given CreateRequest and returns the ID of the newly created company.
// It publishes a "CREATE" event for the new company and returns an error if any operation fails.
func (s *CompanyDataService) CreateCompany(ctx context.Context, request *models.CreateRequest) (string, error) {
	var err error
	id, err := s.repo.CreateCompany(ctx, request)
	if err != nil {
		return "", err
	}
	// Produce a CREATE event to Kafka
	if err = s.producer.Produce(s.topic, id, "CREATE"); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateCompany updates the company information in the repository and publishes an "UPDATE" event to Kafka.
func (s *CompanyDataService) UpdateCompany(ctx context.Context, request *models.UpdateRequest) error {
	var err error
	err = s.repo.UpdateCompany(ctx, request)
	if err != nil {
		return err
	}
	// Produce an UPDATE event to Kafka
	if err = s.producer.Produce(s.topic, request.ID, "UPDATE"); err != nil {
		return err
	}

	return nil
}

// DeleteCompany deletes a company by its ID.
// It first removes the company from the repository and then produces a deletion event to Kafka.
func (s *CompanyDataService) DeleteCompany(ctx context.Context, id string) error {
	var err error
	err = s.repo.DeleteCompany(ctx, id)
	if err != nil {
		return err
	}
	// Produce a DELETE event to Kafka
	if err = s.producer.Produce(s.topic, id, "DELETE"); err != nil {
		return err
	}

	return nil
}
