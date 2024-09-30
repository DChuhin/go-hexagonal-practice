package inputports

import (
	"errors"
	"go-hexagonal-practice/internal/domain"
)

type CreateIoTApplicationCommand struct {
	ApplicationName string
}

type UpdateIoTApplicationCommand struct {
	ApplicationId   string
	ApplicationName string
}

var (
	ErrNotFound = errors.New("application not found")
)

type IoTApplicationService interface {
	CreateApplication(authentication *Authentication, command *CreateIoTApplicationCommand) (*domain.IoTApplication, error)
	UpdateApplication(authentication *Authentication, command *UpdateIoTApplicationCommand) (*domain.IoTApplication, error)
	ListUserApplications(authentication *Authentication) ([]*domain.IoTApplication, error)
	GetApplication(authentication *Authentication, applicationId string) (*domain.IoTApplication, error)
	DeleteApplication(authentication *Authentication, applicationId string) error
}
