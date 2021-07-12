package repository

import (
	"github.com/jmoiron/sqlx"
)

// CompanyPhonePostgres - репозиторий
type CompanyPhonePostgres struct {
	db *sqlx.DB
}

// NewCompanyPhonePostgres - конструктор объекта репозитория
func NewCompanyPhonePostgres(db *sqlx.DB) *CompanyPhonePostgres {
	return &CompanyPhonePostgres{db: db}
}
