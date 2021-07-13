package repository

import (
	"fmt"
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
	query := fmt.Sprintf("INSERT INTO %s (company_id, rubric_id) values ($1, $2)", companyRubricTable)
	_, err := c.db.Exec(query, companyId, rubricId)
	if err != nil {
		return err
	}
	return nil
}
