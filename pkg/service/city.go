package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// CityService - сервис города
type CityService struct {
	repo repository.City
}

// NewCityService - конструктор
func NewCityService(repo repository.City) *CityService {
	return &CityService{repo: repo}
}

// CreateIfNotExists - создать, если не существует
func (c *CityService) CreateIfNotExists(name string) (int, error) {
	return c.repo.CreateIfNotExists(name)
}
