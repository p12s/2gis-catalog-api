package repository

import (
	"github.com/jmoiron/sqlx"
	common "github.com/p12s/2gis-catalog-api"
)

// RubricPostgres - репозиторий
type RubricPostgres struct {
	db *sqlx.DB
}

// NewRubricPostgres - конструктор объекта репозитория
func NewRubricPostgres(db *sqlx.DB) *RubricPostgres {
	return &RubricPostgres{db: db}
}

func (r *RubricPostgres) Create(rubric common.Rubric) (int, error) {
	return 99, nil
}

func (r *RubricPostgres) GetAll(rubricId int) ([]common.Rubric, error) {
	return []common.Rubric{}, nil
}

func (r *RubricPostgres) GetById(rubricId int) (common.Rubric, error) {
	return common.Rubric{}, nil
}

func (r *RubricPostgres) Delete(rubricId int) error {
	return nil
}
