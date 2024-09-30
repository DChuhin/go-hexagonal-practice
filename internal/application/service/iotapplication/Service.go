package iotapplication

import (
	inputports "go-hexagonal-practice/internal/application/ports/input"
	outputports "go-hexagonal-practice/internal/application/ports/output"
	"go-hexagonal-practice/internal/domain"
)

type Service struct {
	iotApplicationStorage outputports.IoTApplicationStorage
}

func NewService(iotApplicationStorage outputports.IoTApplicationStorage) *Service {
	return &Service{
		iotApplicationStorage: iotApplicationStorage,
	}
}

func (s *Service) CreateApplication(authentication *inputports.Authentication, command *inputports.CreateIoTApplicationCommand) (*domain.IoTApplication, error) {
	userId := authentication.UserId
	app, err := domain.NewIoTApplication(userId, command.ApplicationName)
	if err != nil {
		return nil, err
	}
	err = s.iotApplicationStorage.Save(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (s *Service) UpdateApplication(authentication *inputports.Authentication, command *inputports.UpdateIoTApplicationCommand) (*domain.IoTApplication, error) {
	app, err := s.iotApplicationStorage.Get(command.ApplicationId)
	if err != nil {
		return nil, err
	}
	if app.UserId != authentication.UserId {
		return nil, inputports.ErrNotFound
	}
	app.Name = command.ApplicationName
	err = s.iotApplicationStorage.Save(app)
	if err != nil {
		return nil, err
	}
	return app, nil

}

func (s *Service) ListUserApplications(authentication *inputports.Authentication) ([]*domain.IoTApplication, error) {
	apps, err := s.iotApplicationStorage.ListUserApplications(authentication.UserId)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (s *Service) GetApplication(authentication *inputports.Authentication, applicationId string) (*domain.IoTApplication, error) {
	app, err := s.iotApplicationStorage.Get(applicationId)
	if err != nil {
		return nil, err
	}
	if app.UserId != authentication.UserId {
		return nil, inputports.ErrNotFound
	}
	return app, nil
}

func (s *Service) DeleteApplication(authentication *inputports.Authentication, applicationId string) error {
	app, err := s.iotApplicationStorage.Get(applicationId)
	if err != nil {
		return err
	}
	if app.UserId != authentication.UserId {
		return inputports.ErrNotFound
	}
	return s.iotApplicationStorage.Delete(applicationId)
}
