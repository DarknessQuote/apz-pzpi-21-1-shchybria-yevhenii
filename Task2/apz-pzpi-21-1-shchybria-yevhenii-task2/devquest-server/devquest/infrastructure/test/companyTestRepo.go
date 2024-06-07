package test

import (
	"devquest-server/devquest/domain/entities"
	"devquest-server/devquest/domain/repositories"

	"github.com/google/uuid"
)

type companyTestRepo struct{}

func NewCompanyTestRepo() repositories.CompanyRepo {
	return &companyTestRepo{}
}

func (c *companyTestRepo) GetAllCompanies() ([]*entities.Company, error) {
	return []*entities.Company {
		{
			ID: uuid.New(),
			Name: "Company 1",
			Owner: "Owner of Company 1",
			Email: "company1@example.com",
		},
		{
			ID: uuid.New(),
			Name: "Company 2",
			Owner: "Owner of Company 2",
			Email: "company2@example.com",
		},
		{
			ID: uuid.New(),
			Name: "Company 3",
			Owner: "Owner of Company 3",
			Email: "company3@example.com",
		},
	}, nil
}

func (c *companyTestRepo) GetCompanyByID(companyID uuid.UUID) (*entities.Company, error) {
	return &entities.Company{
		ID: companyID,
		Name: "Company by ID",
		Owner: "Owner of Company",
		Email: "company@example.com",
	}, nil
}

func (c *companyTestRepo) AddCompany(company entities.Company) (*entities.Company, error) {
	return &entities.Company{
		ID: uuid.New(),
		Name: "Created Company",
		Owner: "Owner of Created Company",
		Email: "createdcompany@example.com",
	}, nil
}

func (c *companyTestRepo) UpdateCompany(companyID uuid.UUID) error {
	return nil
}

func (c *companyTestRepo) DeleteCompany(companyID uuid.UUID) error {
	return nil
}