package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

type PhoneService struct {
	repo repository.Phone
}

func NewPhoneService(repo repository.Phone) *PhoneService {
	return &PhoneService{repo: repo}
}

func (p *PhoneService) Create(phone common.Phone) (int, error) {
	return p.repo.Create(phone)
}

func (p *PhoneService) GetById(phoneId int) (common.Phone, error) {
	return p.repo.GetById(phoneId)
}

func (p *PhoneService) Delete(phoneId int) error {
	return p.repo.Delete(phoneId)
}
