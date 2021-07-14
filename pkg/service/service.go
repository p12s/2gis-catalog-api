package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// Rubric - рубрика
type Rubric interface {
	Create(name string, parentRubricId int) (int, error)
}

// City - город
type City interface {
	CreateIfNotExists(name string) (int, error)
}

// Street - улица
type Street interface {
	CreateIfNotExists(cityId int, name string) (int, error)
}

// Phone - телефон
type Phone interface {
	Create(companyId int, number string) error
	GetByCompanyId(companyId int) ([]common.Phone, error)
}

// Building - здание
type Building interface {
	Create(cityId, streetId, house int, point string) (int, error)
	CreateNew(building common.BuildingCreateRequest) (int, error)
	GetAllCompany(cityId, streetId, house int) ([]common.CompanyResponse, error)
}

// Company - компания
type Company interface {
	Create(company common.CompanyCreateRequest) (int, error)
	GetById(companyId int) (common.CompanyResponse, error)
	GetAllInBuilding(buildingId int) ([]common.Company, error)
}

// CompanyRubric - рубрика компании
type CompanyRubric interface {
	Create(companyId, rubricId int) error
	GetAllRubricCompany(rubricId int) ([]common.CompanyResponse, error)
}

// Service - сервис
type Service struct {
	Rubric
	City
	Street
	Building
	Company
	CompanyRubric
	Phone
}

// NewService - конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Rubric:        NewRubricService(repos.Rubric),
		City:          NewCityService(repos.City),
		Street:        NewStreetService(repos.Street),
		Building:      NewBuildingService(repos.Building, repos.Company, repos.Phone, repos.City, repos.Street),
		Company:       NewCompanyService(repos.Company, repos.Building, repos.Phone, repos.Rubric, repos.CompanyRubric),
		CompanyRubric: NewCompanyRubricService(repos.CompanyRubric, repos.Phone),
		Phone:         NewPhoneService(repos.Phone),
	}
}
