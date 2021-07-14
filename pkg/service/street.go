package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// StreetService - сревис рубрики
type StreetService struct {
	repo repository.Street
}

// NewStreetService - конструктор
func NewStreetService(repo repository.Street) *StreetService {
	return &StreetService{repo: repo}
}

// CreateIfNotExists - создать, если еще не существует
func (s *StreetService) CreateIfNotExists(cityId int, name string) (int, error) {
	return s.repo.CreateIfNotExists(cityId, name)
}
