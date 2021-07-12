package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type CityService struct {
	repo repository.City
}

func NewCityService(repo repository.City) *CityService {
	return &CityService{repo: repo}
}

func (c *CityService) Create(city common.City) (int, error) {
	return c.repo.Create(city)
}

func (c *CityService) GetById(cityId int) (common.City, error) {
	return c.repo.GetById(cityId)
}

func (c *CityService) Delete(streetId int) error {
	return c.repo.Delete(streetId)
}
