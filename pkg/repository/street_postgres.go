package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

// StreetPostgres - репозиторий
type StreetPostgres struct {
	db *sqlx.DB
}

// NewStreetPostgres - конструктор объекта репозитория
func NewStreetPostgres(db *sqlx.DB) *StreetPostgres {
	return &StreetPostgres{db: db}
}

// CreateIfNotExists - создание, если еще не существует
func (s *StreetPostgres) CreateIfNotExists(cityId int, name string) (int, error) {
	var streetId int
	var isExists bool

	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE city_id=$1 AND name=$2)", streetTable)
	err := s.db.QueryRow(query, cityId, name).Scan(&isExists)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if !isExists {
		query := fmt.Sprintf("INSERT INTO %s (city_id, name) values ($1, $2) RETURNING id", streetTable)
		if err := s.db.QueryRow(query, cityId, name).Scan(&streetId); err != nil {
			return 0, err
		}
	} else {
		query := fmt.Sprintf("SELECT id FROM %s WHERE city_id=$1 AND name=$2", streetTable)
		if err := s.db.QueryRow(query, cityId, name).Scan(&streetId); err != nil {
			return 0, err
		}
	}

	return streetId, nil
}
