package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// CompanyPostgres - репозиторий
type CompanyPostgres struct {
	db *sqlx.DB
	tx *sql.Tx
}

// NewCompanyPostgres - конструктор объекта репозитория
func NewCompanyPostgres(db *sqlx.DB) *CompanyPostgres {
	return &CompanyPostgres{db: db, tx: &sql.Tx{}}
}

// Create - создание компании
// Считаем, что все города и улицы уже занесены в БД, здесь приходят их id, которые валидируем
/* Пришедшее содержимое запроса в объекте common.Company:
{
    "name": "Rambler", // создаем компанию (post)
    "phones": [
        {
            "company_id": 21, // игнорируем
            "number": "32-32-32" // записываем номера в конце - когда имеем id компании
        },
        {
            "company_id": null,
            "number": "33-33-33"
        }
    ],
    "building": {
        "id": null, // игнорируем
        "city_id": 1, // если по совокупности полей city_id  street_id street_id нет записи - создаем, иначе берем id готового
        "street_id": 2,
        "house": 12,
        "point": "(1.234567890, -90.87654321)" // записываем, если пришло (не проверяем)
    },
    "rubric": [
        {
            "id": null, // id всегда есть и он существует (не проверяем)
            "name": "Новая рубрика",
            "parent_rubric_id": null
        }
    ]
}
*/
func (c *CompanyPostgres) Create(company common.CompanyCreateRequest) (int, error) {
	var companyId int
	query := fmt.Sprintf("INSERT INTO %s (name, building_id) values ($1, $2) RETURNING id", companyTable)
	row := c.db.QueryRow(query, company.Name, company.Building.Id)
	if err := row.Scan(&companyId); err != nil {
		return 0, err
	}
	return companyId, nil
}

func (c *CompanyPostgres) Exists(name string) (bool, error) {
	var isExists bool
	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE name=$1 LIMIT 1)", companyTable)
	err := c.db.QueryRow(query, name).Scan(&isExists)
	if err != nil && err != sql.ErrNoRows {
		return isExists, err
	}
	return isExists, nil
}

func (c *CompanyPostgres) OpenTransaction() error {
	var err error
	c.tx, err = c.db.Begin()
	if err != nil {
		return err
	}

	return nil
}

func (c *CompanyPostgres) RollbackTransaction() error {
	err := c.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

func (c *CompanyPostgres) CloseTransaction() error {
	err := c.tx.Commit()
	if err != nil {
		return err
	}
	return nil
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
