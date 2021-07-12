package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
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
	Create(phone common.Phone) (int, error) // добавление в справочник здания вместе с организациями, которые в ней находятся
}

// Company - компания
type Company interface {
	Create(company common.Company) (int, error)
	GetById(companyId int) (common.Company, error)             // выдачу информации об организациях по их идентификаторам
	GetAllInBuilding(buildingId int) ([]common.Company, error) // выдачу всех организаций находящихся в конкретном здании
	GetAllByRubricId(rubricId int) ([]common.Company, error)   // выдачу списка всех организаций, которые относятся к указанной рубрике
	Delete(companyId int) error
}

type CompanyRubric interface {
	Create(companyId, rubricId int) error
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
		Rubric:   NewRubricPostgres(db),
		City:     NewCityPostgres(db),
		Street:   NewStreetPostgres(db),
		Building: NewBuildingPostgres(db),
		Company:  NewCompanyPostgres(db),
		CompanyRubric:  NewCompanyRubricPostgres(db),
		Phone:    NewPhonePostgres(db),
	}
}
