package service

import (
	"github.com/p12s/2gis-catalog-api"
	"github.com/p12s/2gis-catalog-api/pkg/repository"
)

// PhoneService - сервис телефона компании
type PhoneService struct {
	repo repository.Phone
}

// NewPhoneService - конструктор
func NewPhoneService(repo repository.Phone) *PhoneService {
	return &PhoneService{repo: repo}
}

// Create - создать
func (p *PhoneService) Create(companyId int, number string) error {
	return p.repo.Create(companyId, number)
}

// GetByCompanyId - получить
func (p *PhoneService) GetByCompanyId(companyId int) ([]common.Phone, error) {
	return p.repo.GetByCompanyId(companyId)
}
