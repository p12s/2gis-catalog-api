package repository

import (
	"fmt"
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

func (p *PhonePostgres) Create(companyId int, number string) error {
	query := fmt.Sprintf("INSERT INTO %s (company_id, number) values ($1, $2)", phoneTable)
	_, err := p.db.Exec(query, companyId, number)
	if err != nil {
		return err
	}
	return nil
}

func (p *PhonePostgres) GetById(phoneId int) (common.Phone, error) {
	return common.Phone{}, nil
}

func (p *PhonePostgres) Delete(phoneId int) error {
	return nil
}
