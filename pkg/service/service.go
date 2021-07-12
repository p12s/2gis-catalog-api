package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type Rubric interface {
	Create(rubric common.Rubric) (int, error)
	GetAll(rubricId int) ([]common.Rubric, error)
	GetById(rubricId int) (common.Rubric, error)
	Delete(rubricId int) error
}

type City interface {
	Create(city common.City) (int, error)
	GetById(cityId int) (common.City, error)
	Delete(streetId int) error
}

type Street interface {
	Create(street common.Street) (int, error)
	GetByName(streetName string, cityId int) (common.Street, error)
	Delete(streetId int) error
}

type Phone interface {
	Create(phone common.Phone) (int, error)
	GetById(phoneId int) (common.Phone, error)
	Delete(phoneId int) error
}

type Building interface {
	Create(phone common.Phone) (int, error)
}

// Company - компания
type Company interface {
	Create(company common.Company) (int, error)
	GetById(companyId int) (common.Company, error)
	GetAllInBuilding(buildingId int) ([]common.Company, error)
	GetAllByRubricId(rubricId int) ([]common.Company, error)
	Delete(companyId int) error
}

type CompanyRubric interface {
	Create(companyId, rubricId int) error
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
		Rubric:   NewRubricService(repos.Rubric),
		City:     NewCityService(repos.City),
		Street:   NewStreetService(repos.Street),
		Building: NewBuildingService(repos.Building),
		Company:  NewCompanyService(repos.Company),
		CompanyRubric:  NewCompanyRubricService(repos.CompanyRubric),
		Phone:    NewPhoneService(repos.Phone),
	}
}
