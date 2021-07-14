package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
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

func (c *CompanyRubricPostgres) GetAllRubricCompany(rubricId int) ([]common.Company, error) {
	var items []common.Company
	query := fmt.Sprintf(`SELECT c.id, c.name FROM %s c 
		INNER JOIN %s cr on cr.company_id = c.id
		WHERE cr.rubric_id = $1`,
		companyTable, companyRubricTable)
	if err := c.db.Select(&items, query, rubricId); err != nil {
		return nil, err
	}

	return items, nil
}
