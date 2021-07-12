package repository

import (
	"github.com/jmoiron/sqlx"
)

// CompanyRubricPostgres - репозиторий
type CompanyRubricPostgres struct {
	db *sqlx.DB
}

// NewCompanyRubricPostgres - конструктор объекта репозитория
func NewCompanyRubricPostgres(db *sqlx.DB) *CompanyRubricPostgres {
	return &CompanyRubricPostgres{db: db}
}

func (c *CompanyRubricPostgres) Create(companyId, rubricId int) error {
	return nil
}