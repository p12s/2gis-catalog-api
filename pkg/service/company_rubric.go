package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CompanyRubricService struct {
	repo repository.CompanyRubric
}

func NewCompanyRubricService(repo repository.CompanyRubric) *CompanyRubricService {
	return &CompanyRubricService{repo: repo}
}

func (c *CompanyRubricService) Create(companyId, rubricId int) error {
	return c.repo.Create(companyId, rubricId)
}
