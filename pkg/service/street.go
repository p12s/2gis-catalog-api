package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type StreetService struct {
	repo repository.Street
}

func NewStreetService(repo repository.Street) *StreetService {
	return &StreetService{repo: repo}
}

func (s *StreetService) Create(street common.Street) (int, error) {
	return s.repo.Create(street)
}

func (s *StreetService) GetByName(streetName string, cityId int) (common.Street, error) {
	return s.repo.GetByName(streetName, cityId)
}

func (s *StreetService) Delete(streetId int) error {
	return s.repo.Delete(streetId)
}
