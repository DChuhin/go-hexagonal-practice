package mongodbiotapplication

import (
	"context"
	"errors"
	inputports "go-hexagonal-practice/internal/application/ports/input"
	"go-hexagonal-practice/internal/domain"
	"go-hexagonal-practice/internal/integration/mongodb/repository"
	"go-hexagonal-practice/internal/util/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	ctx        context.Context
	repository repository.MongoRepository[Document]
}

func NewStorage(ctx context.Context, mongoRepository repository.MongoRepository[Document]) *Storage {
	return &Storage{
		ctx:        ctx,
		repository: mongoRepository,
	}
}

func (s *Storage) Save(app *domain.IoTApplication) error {
	document, err := mapper.Map[Document, domain.IoTApplication](app)
	if err != nil {
		return err
	}
	_, err = s.repository.InsertOne(s.ctx, document)
	return err
}

func (s *Storage) Get(id string) (*domain.IoTApplication, error) {
	document, err := s.repository.FindById(s.ctx, id)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, inputports.ErrNotFound
	}
	app, err := mapper.Map[domain.IoTApplication, Document](document)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (s *Storage) ListUserApplications(userId string) ([]*domain.IoTApplication, error) {
	documents, err := s.repository.Find(s.ctx, bson.D{{"userId", userId}})
	if err != nil {
		return nil, err
	}
	apps, err := mapper.MapSlice[domain.IoTApplication, Document](documents)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (s *Storage) Delete(id string) error {
	_, err := s.repository.DeleteById(s.ctx, id)
	return err
}
