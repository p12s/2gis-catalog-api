package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// RubricPostgres - репозиторий
type RubricPostgres struct {
	db *sqlx.DB
}

// NewRubricPostgres - конструктор объекта репозитория
func NewRubricPostgres(db *sqlx.DB) *RubricPostgres {
	return &RubricPostgres{db: db}
}

// Create - создание
func (r *RubricPostgres) Create(name string, parentRubricId int) (int, error) {
	var row *sql.Row
	if parentRubricId == 0 {
		query := fmt.Sprintf("INSERT INTO %s (name) values ($1) RETURNING id", rubricTable)
		row = r.db.QueryRow(query, name)
	} else {
		query := fmt.Sprintf("INSERT INTO %s (name, parent_rubric_id) values ($1, $2) RETURNING id", rubricTable)
		row = r.db.QueryRow(query, name, parentRubricId)
	}

	var rubricId int
	if err := row.Scan(&rubricId); err != nil {
		return 0, err
	}

	return rubricId, nil
}

// FindByName - поиска по имени
func (r *RubricPostgres) FindByName(rubricName string) (common.Rubric, error) {
	var exists common.Rubric
	query := fmt.Sprintf("SELECT * FROM %s WHERE name=$1", rubricTable)
	err := r.db.Select(&exists, query, rubricName)
	if err != nil {
		return exists, err
	}
	return exists, err
}
