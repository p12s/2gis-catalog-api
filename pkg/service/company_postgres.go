package repository

import (
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/2gis-catalog-api"
)

// CompanyPostgres - репозиторий
type CompanyPostgres struct {
	db *sqlx.DB
}

// NewCompanyPostgres - конструктор объекта репозитория
func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db}
}

func (c *CompanyPostgres) Create(company common.Company) (int, error) {
	return 99, nil
}

func (c *CompanyPostgres) GetById(companyId int) (common.Company, error) {
	return common.Company{}, nil
}

func (c *CompanyPostgres) GetAllInBuilding(buildingId int) ([]common.Company, error) {
	return []common.Company{}, nil
}

func (c *CompanyPostgres) GetAllByRubricId(rubricId int) ([]common.Company, error) {
	return []common.Company{}, nil
}

func (c *CompanyPostgres) Delete(companyId int) error {
	return nil
}
