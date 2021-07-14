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
// Считаем, что все города и улицы уже занесены в БД, в текущем запросе приходят их id, которые не валидируем
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

// Exists - проверка компании на существование
func (c *CompanyPostgres) Exists(name string) (bool, error) {
	var isExists bool
	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE name=$1 LIMIT 1)", companyTable)
	err := c.db.QueryRow(query, name).Scan(&isExists)
	if err != nil && err != sql.ErrNoRows {
		return isExists, err
	}
	return isExists, nil
}

// OpenTransaction - открытие транзакци (нужно тестить, возможно не работает так как нужно)
func (c *CompanyPostgres) OpenTransaction() error {
	var err error
	c.tx, err = c.db.Begin()
	if err != nil {
		return err
	}

	return nil
}

// RollbackTransaction - откат транзакции
func (c *CompanyPostgres) RollbackTransaction() error {
	err := c.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}

// CloseTransaction - закрыте транзакции
func (c *CompanyPostgres) CloseTransaction() error {
	err := c.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// GetById - получить компанию по id
func (c *CompanyPostgres) GetById(companyId int) (common.Company, error) {
	var company common.Company
	query := fmt.Sprintf(`SELECT id, name FROM %s WHERE id = $1`, companyTable)
	if err := c.db.Get(&company, query, companyId); err != nil {
		return company, err
	}

	return company, nil
}

// GetAllInBuilding - получить все компании по зданию
func (c *CompanyPostgres) GetAllInBuilding(buildingId int) ([]common.Company, error) {
	var items []common.Company
	query := fmt.Sprintf(`SELECT id, name FROM %s WHERE building_id = $1`, companyTable)
	if err := c.db.Select(&items, query, buildingId); err != nil {
		return nil, err
	}

	return items, nil
}
