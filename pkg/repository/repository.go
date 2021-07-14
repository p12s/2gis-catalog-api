package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

type Rubric interface {
	Create(name string, parentRubricId int) (int, error)
	GetAll(rubricId int) ([]common.Rubric, error)
	GetById(rubricId int) (common.Rubric, error)
	Delete(rubricId int) error
	FindByName(rubricName string) (common.Rubric, error)
}

type City interface {
	CreateIfNotExists(name string) (int, error)
}

type Street interface {
	CreateIfNotExists(cityId int, name string) (int, error)
}

type Phone interface {
	Create(companyId int, number string) error
	GetByCompanyId(companyId int) ([]common.Phone, error)
}

type Building interface {
	Create(cityId, streetId, house int, point string) (int, error)
	CreateNew(building common.BuildingCreateRequest) (int, error)
	GetByCityStreetHouse(cityId int, streetId int, house int) (common.Building, error)
}

// Company - компания
type Company interface {
	Create(company common.CompanyCreateRequest) (int, error)
	GetById(companyId int) (common.Company, error)
	GetAllInBuilding(buildingId int) ([]common.Company, error)
	GetAllByRubricId(rubricId int) ([]common.Company, error)
	Delete(companyId int) error
	Exists(name string) (bool, error)
	OpenTransaction() error
	RollbackTransaction() error
	CloseTransaction() error
}

type CompanyRubric interface {
	Create(companyId, rubricId int) error
	GetAllRubricCompany(rubricId int) ([]common.Company, error)
}

// Repository - репозитрий
type Repository struct {
	Rubric
	City
	Street
	Building
	Company
	CompanyRubric
	Phone
}

// NewRepository - конструктор
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Rubric:        NewRubricPostgres(db),
		City:          NewCityPostgres(db),
		Street:        NewStreetPostgres(db),
		Building:      NewBuildingPostgres(db),
		Company:       NewCompanyPostgres(db),
		CompanyRubric: NewCompanyRubricPostgres(db),
		Phone:         NewPhonePostgres(db),
	}
}
