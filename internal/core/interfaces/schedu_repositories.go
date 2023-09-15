package interfaces

import "github.com/HEEPOKE/go-echo-hexagonal-api/internal/domains/models"

type ScheduRepository interface {
	List() ([]*models.Schedu, error)
	CreateNoti(schedu *models.Schedu) error
}
