package service

import (
	common "github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type BuildingService struct {
	repo        repository.Building
	repoCompany repository.Company
	repoPhone   repository.Phone
}

func NewBuildingService(
	repo repository.Building,
	repoCompany repository.Company,
	repoPhone repository.Phone,
) *BuildingService {
	return &BuildingService{
		repo:        repo,
		repoCompany: repoCompany,
		repoPhone:   repoPhone,
	}
}

func (b *BuildingService) Create(cityId, streetId, house int, point string) (int, error) {
	return b.repo.Create(cityId, streetId, house, point)
}

func (b *BuildingService) GetAllCompany(cityId, streetId, house int) ([]common.CompanyResponse, error) {

	var companies []common.CompanyResponse

	existsBuilding, err := b.repo.GetByCityStreetHouse(cityId, streetId, house)
	if err != nil || existsBuilding.Id == 0 {
		return companies, err
	}

	companyList, err := b.repoCompany.GetAllInBuilding(existsBuilding.Id)
	if err != nil {
		return companies, err
	}

	for _, company := range companyList {
		// телефоны
		companyPhones, err := b.repoPhone.GetByCompanyId(company.Id)
		if err != nil {
			return companies, err
		}
		// добавляем к компании
		companies = append(companies, common.CompanyResponse{
			Name:   company.Name,
			Phones: companyPhones,
		})
	}

	return companies, nil
}
