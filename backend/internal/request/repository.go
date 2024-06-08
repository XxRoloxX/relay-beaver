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
	Create(request models.ProxiedRequest) (models.ProxiedRequest, error)
	FindById(id string) (models.ProxiedRequest, error)
	FindAll() ([]models.ProxiedRequest, error)
	Update(id string, request models.ProxiedRequest) error
	Delete(id string) error
}

type RequestMongoRepository struct {
	Db mongo.Database
}

func (repo RequestMongoRepository) getRequestsCollection() *mongo.Collection {
	return repo.Db.Collection(RequestCollection)
}

func (repo RequestMongoRepository) Create(request models.ProxiedRequest) (models.ProxiedRequest, error) {
	res, error := repo.getRequestsCollection().InsertOne(context.TODO(), request)

	if error != nil {
		return models.ProxiedRequest{}, error
	}

	request.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return request, nil
}

func (repo RequestMongoRepository) FindById(id string) (models.ProxiedRequest, error) {
	requestId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return models.ProxiedRequest{}, err
	}

	res := repo.getRequestsCollection().FindOne(context.TODO(), bson.M{"_id": requestId})

	if res.Err() != nil {
		return models.ProxiedRequest{}, res.Err()
	}

	var request models.ProxiedRequest
	err = res.Decode(&request)

	return request, err
}

func (repo RequestMongoRepository) FindAll() ([]models.ProxiedRequest, error) {
	res, err := repo.getRequestsCollection().Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var requests []models.ProxiedRequest
	for res.Next(context.Background()) {
		var request models.ProxiedRequest
		err := res.Decode(&request)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func (repo RequestMongoRepository) Update(id string, request models.ProxiedRequest) error {
	return nil
}

func (repo RequestMongoRepository) Delete(id string) error {
	return nil
}
