package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type RubricService struct {
	repo repository.Rubric
}

func NewRubricService(repo repository.Rubric) *RubricService {
	return &RubricService{repo: repo}
}

func (r *RubricService) Create(name string, parentRubricId int) (int, error) {
	return r.repo.Create(name, parentRubricId)
}

func (r *RubricService) GetAll(rubricId int) ([]common.Rubric, error) {
	return r.repo.GetAll(rubricId)
}

func (r *RubricService) GetById(rubricId int) (common.Rubric, error) {
	return r.repo.GetById(rubricId)
}

func (r *RubricService) Delete(rubricId int) error {
	return r.repo.Delete(rubricId)
}
