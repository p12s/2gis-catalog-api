package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// Rubric - рубрика (предзаполненные в БД - Транспорт, Недвижимость и т.п.)
type Rubric interface {
	Create(name string, parentRubricId int) (int, error)
	FindByName(rubricName string) (common.Rubric, error)
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
	GetByCityStreetHouse(cityId int, streetId int, house int) (common.Building, error)
}

// Company - компания
type Company interface {
	Create(company common.CompanyCreateRequest) (int, error)
	GetById(companyId int) (common.Company, error)
	GetAllInBuilding(buildingId int) ([]common.Company, error)
	Exists(name string) (bool, error)
	OpenTransaction() error
	RollbackTransaction() error
	CloseTransaction() error
}

// CompanyRubric - рубрика компании
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
