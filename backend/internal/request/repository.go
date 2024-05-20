package request

import (
	"backend/pkg/models"
	UUID "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestRepository interface {
	Store(request models.Request) error
	FindById(id UUID.UUID) (models.Request, error)
	FindAll() ([]models.Request, error)
	Update(request models.Request) error
	Delete(id UUID.UUID) error
}
type RequestMongoRepository struct {
	Db mongo.Database
}

func (repo RequestMongoRepository) Store(request models.Request) error {
	return nil
}

func (repo RequestMongoRepository) FindById(id UUID.UUID) (models.Request, error) {
	return models.Request{}, nil
}

func (repo RequestMongoRepository) FindAll() ([]models.Request, error) {
	mockRequest := models.Request{
		Id:     UUID.New(),
		Source: "source",
		Destination: models.Address{
			Host: "host",
			Port: 8080,
		},
		StartTimestamp:  0,
		FinishTimestamp: 0,
		Headers: map[string]string{
			"header": "exampleHeader",
		},
		Body:         "body",
		Method:       "GET",
		ResponseCode: 200,
	}

	return []models.Request{mockRequest}, nil
}

func (repo RequestMongoRepository) Update(request models.Request) error {
	return nil
}

func (repo RequestMongoRepository) Delete(id UUID.UUID) error {
	return nil
}
