package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CompanyService struct {
	repo repository.Company
}

func NewCompanyService(repo repository.Company) *CompanyService {
	return &CompanyService{repo: repo}
}

func (c *CompanyService) Create(company common.Company) (int, error) {
	return c.repo.Create(company)
}

func (c *CompanyService) GetById(companyId int) (common.Company, error) {
	return c.repo.GetById(companyId)
}

func (c *CompanyService) GetAllInBuilding(buildingId int) ([]common.Company, error) {
	return c.repo.GetAllInBuilding(buildingId)
}

func (c *CompanyService) GetAllByRubricId(rubricId int) ([]common.Company, error) {
	return c.repo.GetAllByRubricId(rubricId)
}

func (c *CompanyService) Delete(companyId int) error {
	return c.repo.Delete(companyId)
}