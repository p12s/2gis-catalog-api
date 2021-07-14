package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CityService struct {
	repo repository.City
}

func NewCityService(repo repository.City) *CityService {
	return &CityService{repo: repo}
}

func (c *CityService) CreateIfNotExists(name string) (int, error) {
	return c.repo.CreateIfNotExists(name)
}
