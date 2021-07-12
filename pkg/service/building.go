package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type BuildingService struct {
	repo repository.Building
}

func NewBuildingService(repo repository.Building) *BuildingService {
	return &BuildingService{repo: repo}
}

func (b *BuildingService) Create(phone common.Phone) (int, error) {
	return b.repo.Create(phone)
}
