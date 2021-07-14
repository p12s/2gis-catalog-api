package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// CityPostgres - репозиторий
type CityPostgres struct {
	db *sqlx.DB
}

// NewCityPostgres - конструктор объекта репозитория
func NewCityPostgres(db *sqlx.DB) *CityPostgres {
	return &CityPostgres{db: db}
}

// CreateIfNotExists - создание города, если он еще не существует
func (c *CityPostgres) CreateIfNotExists(name string) (int, error) {
	var cityId int
	var isExists bool

	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE name=$1)", cityTable)
	err := c.db.QueryRow(query, name).Scan(&isExists)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if !isExists {
		query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", cityTable)
		row := c.db.QueryRow(query, name)
		if err := row.Scan(&cityId); err != nil {
			return 0, err
		}
	} else {
		query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", cityTable)
		err := c.db.QueryRow(query, name).Scan(&cityId)
		if err != nil {
			return 0, err
		}
	}

	return cityId, nil
}
