package service

import (
	common "github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// BuildingService - сервис здания
type BuildingService struct {
	repo        repository.Building
	repoCompany repository.Company
	repoPhone   repository.Phone
	repoCity    repository.City
	repoStreet  repository.Street
}

// NewBuildingService - конструктор
func NewBuildingService(
	repo repository.Building,
	repoCompany repository.Company,
	repoPhone repository.Phone,
	repoCity repository.City,
	repoStreet repository.Street,
) *BuildingService {
	return &BuildingService{
		repo:        repo,
		repoCompany: repoCompany,
		repoPhone:   repoPhone,
		repoCity:    repoCity,
		repoStreet:  repoStreet,
	}
}

// Create - запись адреса здания по существующим в БД id города, улицы
func (b *BuildingService) Create(cityId, streetId, house int, point string) (int, error) {
	return b.repo.Create(cityId, streetId, house, point)
}

// CreateNew - запись нового адреса, по еще не добавленнным в БД городу и улице
func (b *BuildingService) CreateNew(building common.BuildingCreateRequest) (int, error) {
	err := b.repoCompany.OpenTransaction()
	if err != nil {
		return 0, err
	}

	// сохраним город, если пришло только его название
	if building.CityId == 0 && building.City != "" {
		cityId, err := b.repoCity.CreateIfNotExists(building.City)
		if err != nil {

			rollbackErr := b.repoCompany.RollbackTransaction()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}
		building.CityId = cityId
	}

	// сохраним улицу, если пришло только ее название
	if building.StreetId == 0 && building.Street != "" {
		streetId, err := b.repoStreet.CreateIfNotExists(building.CityId, building.Street)
		if err != nil {

			rollbackErr := b.repoCompany.RollbackTransaction()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}
		building.StreetId = streetId
	}

	buildingId, err := b.repo.CreateNew(building)
	if err != nil {
		rollbackErr := b.repoCompany.RollbackTransaction()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	return buildingId, b.repoCompany.CloseTransaction()
}

// GetAllCompany - получить все компании по адресу здания
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
