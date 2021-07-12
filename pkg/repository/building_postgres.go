package repository

import (
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/2gis-catalog-api"
)

// BuildingPostgres - репозиторий
type BuildingPostgres struct {
	db *sqlx.DB
}

// NewBuildingPostgres - конструктор объекта репозитория
func NewBuildingPostgres(db *sqlx.DB) *BuildingPostgres {
	return &BuildingPostgres{db: db}
}

func (b *BuildingPostgres) Create(phone common.Phone) (int, error) {
	return 99, nil
}
