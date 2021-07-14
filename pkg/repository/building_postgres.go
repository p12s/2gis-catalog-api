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

// Create - добаление адреса здания по существующим id города и улицы
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

// CreateNew - создает новое здание по новому адресу
// город и улица могт быть новыми - создаем полностью новый адрес
func (b *BuildingPostgres) CreateNew(building common.BuildingCreateRequest) (int, error) {

	exists, err := b.GetByCityStreetHouse(building.CityId, building.StreetId, building.House)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if exists.Id == 0 {
		query := fmt.Sprintf("INSERT INTO %s (city_id, street_id, house, point) values ($1, $2, $3, $4) RETURNING id", buildingTable)
		row := b.db.QueryRow(query, building.CityId, building.StreetId, building.House, building.Point)
		if err := row.Scan(&exists.Id); err != nil {
			return 0, err
		}
	}

	return exists.Id, nil
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
