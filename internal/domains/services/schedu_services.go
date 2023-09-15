package services

import (
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/core/interfaces"
	"github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"
)

type ScheduService struct {
	scheduRepository interfaces.ScheduRepository
}

func NewScheduService(scheduRepository interfaces.ScheduRepository) *ScheduService {
	return &ScheduService{scheduRepository: scheduRepository}
}

func (s *ScheduService) List() ([]*models.Schedu, error) {
	return s.scheduRepository.List()
}

func (s *ScheduService) CreateNoti(schedu *models.Schedu) error {
	return s.scheduRepository.CreateNoti(schedu)
}
