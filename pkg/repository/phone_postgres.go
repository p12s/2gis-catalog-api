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

func (p *PhonePostgres) GetByCompanyId(companyId int) ([]common.Phone, error) {
	var items []common.Phone
	query := fmt.Sprintf(`SELECT * FROM %s WHERE company_id = $1`, phoneTable)
	if err := p.db.Select(&items, query, companyId); err != nil {
		return nil, err
	}
	return items, nil
}
