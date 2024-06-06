package repositories

import (
	"devquest-server/devquest/domain/models"

	"github.com/google/uuid"
)

type CompanyRepo interface {
	GetAllCompanies() ([]*models.Company, error)
	GetCompanyByID(uuid.UUID) (*models.Company, error)
	AddCompany(models.Company) (*models.Company, error)
	UpdateCompany(uuid.UUID) error
	DeleteCompany(uuid.UUID) error
}