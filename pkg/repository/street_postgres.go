package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// StreetPostgres - репозиторий
type StreetPostgres struct {
	db *sqlx.DB
}

// NewStreetPostgres - конструктор объекта репозитория
func NewStreetPostgres(db *sqlx.DB) *StreetPostgres {
	return &StreetPostgres{db: db}
}

func (s *StreetPostgres) Create(street common.Street) (int, error) {
	return 99, nil
}

func (s *StreetPostgres) GetByName(streetName string, cityId int) (common.Street, error) {
	return common.Street{}, nil
}

func (s *StreetPostgres) Delete(streetId int) error {
	return nil
}
