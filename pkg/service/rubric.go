package service

import (
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// RubricService - сервис рубрики
type RubricService struct {
	repo repository.Rubric
}

// NewRubricService - конструктор
func NewRubricService(repo repository.Rubric) *RubricService {
	return &RubricService{repo: repo}
}

// Create - создать
func (r *RubricService) Create(name string, parentRubricId int) (int, error) {
	return r.repo.Create(name, parentRubricId)
}
