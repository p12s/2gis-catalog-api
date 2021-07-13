package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// BuildingPostgres - репозиторий
type BuildingPostgres struct {
	db *sqlx.DB
}

// NewBuildingPostgres - конструктор объекта репозитория
func NewBuildingPostgres(db *sqlx.DB) *BuildingPostgres {
	return &BuildingPostgres{db: db}
}

// Create - добаление здания
// считаем, что принятые id города и улицы существуют в рамках БД (добавлены ранее)
func (b *BuildingPostgres) Create(cityId, streetId, house int, point string) (int, error) {
	var buildingId int
	query := fmt.Sprintf("INSERT INTO %s (city_id, street_id, house, point) values ($1, $2, $3, $4) RETURNING id", buildingTable)

	row := b.db.QueryRow(query, cityId, streetId, house, point)
	if err := row.Scan(&buildingId); err != nil {
		return 0, err
	}
	return buildingId, nil
}

func (b *BuildingPostgres) GetByCityStreetHouse(cityId int, streetId int, house int) (common.Building, error) {
	var exists common.Building
	var isExists bool

	query := fmt.Sprintf("SELECT exists (SELECT * FROM %s WHERE city_id=$1 AND street_id=$2 AND house=$3)", buildingTable)
	err := b.db.QueryRow(query, cityId, streetId, house).Scan(&isExists)
	if err != nil && err != sql.ErrNoRows {
		return exists, err
	}

	if isExists {
		query := fmt.Sprintf("SELECT * FROM %s WHERE city_id=$1 AND street_id=$2 AND house=$3", buildingTable)
		err := b.db.Get(&exists, query, cityId, streetId, house)
		if err != nil {
			return exists, err
		}
	}

	return exists, err
}
