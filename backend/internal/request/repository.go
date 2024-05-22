package request

import (
	"backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	RequestCollection = "requests"
)

type RequestRepository interface {
	Create(request models.Request) error
	FindById(id string) (models.Request, error)
	FindAll() ([]models.Request, error)
	Update(id string, request models.Request) error
	Delete(id string) error
}

type RequestMongoRepository struct {
	Db mongo.Database
}

func (repo RequestMongoRepository) getRequestsCollection() *mongo.Collection {
	return repo.Db.Collection(RequestCollection)
}

func (repo RequestMongoRepository) Create(request models.Request) error {
	_, error := repo.getRequestsCollection().InsertOne(context.TODO(), request)

	return error
}

func (repo RequestMongoRepository) FindById(id string) (models.Request, error) {
	requestId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.Request{}, err
	}

	res := repo.getRequestsCollection().FindOne(context.TODO(), bson.M{"_id": requestId})

	if res.Err() != nil {
		return models.Request{}, res.Err()
	}

	var request models.Request
	err = res.Decode(&request)

	return request, err
}

func (repo RequestMongoRepository) FindAll() ([]models.Request, error) {
	res, err := repo.getRequestsCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var requests []models.Request
	for res.Next(context.Background()) {
		var request models.Request
		err := res.Decode(&request)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func (repo RequestMongoRepository) Update(id string, request models.Request) error {
	return nil
}

func (repo RequestMongoRepository) Delete(id string) error {
	return nil
}
