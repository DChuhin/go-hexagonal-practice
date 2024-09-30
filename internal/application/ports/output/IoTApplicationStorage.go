package output_ports

import "go-hexagonal-practice/internal/domain"

type IoTApplicationStorage interface {
	Save(app *domain.IoTApplication) error
	Get(id string) (*domain.IoTApplication, error)
	ListUserApplications(userId string) ([]*domain.IoTApplication, error)
	Delete(id string) error
}
