package service

import (
	common "github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CompanyRubricService struct {
	repo      repository.CompanyRubric
	repoPhone repository.Phone
}

func NewCompanyRubricService(repo repository.CompanyRubric, repoPhone repository.Phone) *CompanyRubricService {
	return &CompanyRubricService{repo: repo, repoPhone: repoPhone}
}

func (c *CompanyRubricService) Create(companyId, rubricId int) error {
	return c.repo.Create(companyId, rubricId)
}

func (c *CompanyRubricService) GetAllRubricCompany(rubricId int) ([]common.CompanyResponse, error) {
	var items []common.CompanyResponse
	companies, err := c.repo.GetAllRubricCompany(rubricId)
	if err != nil {
		return items, err
	}

	for _, company := range companies {
		phones, err := c.repoPhone.GetByCompanyId(company.Id)
		if err != nil {
			return items, err
		}
		items = append(items, common.CompanyResponse{
			Name:   company.Name,
			Phones: phones,
		})
	}

	return items, nil
}
