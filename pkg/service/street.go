package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type StreetService struct {
	repo repository.Street
}

func NewStreetService(repo repository.Street) *StreetService {
	return &StreetService{repo: repo}
}

func (s *StreetService) CreateIfNotExists(cityId int, name string) (int, error) {
	return s.repo.CreateIfNotExists(cityId, name)
}
