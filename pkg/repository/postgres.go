package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	rubricTable        = "rubric"
	cityTable          = "city"
	streetTable        = "street"
	buildingTable      = "building"
	companyTable       = "company"
	companyRubricTable = "company_rubric"
	phoneTable         = "phone"
)

// Config - конфиг БД
type Config struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB - коннект к БД
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
