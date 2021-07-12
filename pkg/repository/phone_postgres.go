package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/2gis-catalog-api"
)

// PhonePostgres - репозиторий
type PhonePostgres struct {
	db *sqlx.DB
}

// NewPhonePostgres - конструктор объекта репозитория
func NewPhonePostgres(db *sqlx.DB) *PhonePostgres {
	return &PhonePostgres{db: db}
}

func (p *PhonePostgres) Create(phone common.Phone) (int, error) {
	return 99, nil
}

func (p *PhonePostgres) GetById(phoneId int) (common.Phone, error) {
	return common.Phone{}, nil
}

func (p *PhonePostgres) Delete(phoneId int) error {
	return nil
}
