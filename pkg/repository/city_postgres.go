package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// CityPostgres - репозиторий
type CityPostgres struct {
	db *sqlx.DB
}

// NewCityPostgres - конструктор объекта репозитория
func NewCityPostgres(db *sqlx.DB) *CityPostgres {
	return &CityPostgres{db: db}
}

func (c *CityPostgres) Create(city common.City) (int, error) {
	return 99, nil
}

func (c *CityPostgres) GetById(cityId int) (common.City, error) {
	return common.City{}, nil
}

func (c *CityPostgres) Delete(streetId int) error {
	return nil
}
