package loadbalancer

import (
	"backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	LoadBalancerCollection = "load_balancers"
)

type LoadBalancerMongoRepository struct {
	Db mongo.Database
}

func (lb LoadBalancerMongoRepository) getLoadBalancersCollection() *mongo.Collection {
	return lb.Db.Collection(LoadBalancerCollection)
}

type LoadBalancerRepository interface {
	Create(lb models.LoadBalancer) (models.LoadBalancer, error)
	FindById(id string) (models.LoadBalancer, error)
	FindAll() ([]models.LoadBalancer, error)
	Delete(id string) error
	Update(id string, lb models.LoadBalancer) (models.LoadBalancer, error)
}

func (l LoadBalancerMongoRepository) FindById(id string) (models.LoadBalancer, error) {

	primitiveId, idError := primitive.ObjectIDFromHex(id)

	if idError != nil {
		return models.LoadBalancer{}, idError
	}

	res := l.getLoadBalancersCollection().FindOne(context.TODO(), bson.M{"_id": primitiveId})

	if res.Err() != nil {
		return models.LoadBalancer{}, res.Err()
	}

	var lb models.LoadBalancer

	err := res.Decode(&lb)

	if err != nil {
		return models.LoadBalancer{}, err
	}

	return lb, nil
}

func (l LoadBalancerMongoRepository) Create(lb models.LoadBalancer) (models.LoadBalancer, error) {
	res, err := l.getLoadBalancersCollection().InsertOne(context.TODO(), lb)

	if err != nil {
		return models.LoadBalancer{}, err
	}

	lb.Id = res.InsertedID.(primitive.ObjectID).Hex()
	return lb, nil
}

func (l LoadBalancerMongoRepository) FindAll() ([]models.LoadBalancer, error) {
	entries, err := l.getLoadBalancersCollection().Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var lbs []models.LoadBalancer
	for entries.Next(context.TODO()) {
		var lb = models.LoadBalancer{}
		err = entries.Decode(&lb)
		if err != nil {
			return nil, err
		}
		lbs = append(lbs, lb)

	}

	return lbs, nil
}

func (l LoadBalancerMongoRepository) Delete(id string) error {
	primitiveId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = l.getLoadBalancersCollection().DeleteOne(context.TODO(), bson.M{"_id": primitiveId})
	return err
}

func (l LoadBalancerMongoRepository) Update(id string, lb models.LoadBalancer) (models.LoadBalancer, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.LoadBalancer{}, err
	}
	filter := bson.D{{"_id", objID}}

	updateDoc := bson.M{}
	bsonBytes, _ := bson.Marshal(&lb)
	err = bson.Unmarshal(bsonBytes, &updateDoc)
	if err != nil {
		return models.LoadBalancer{}, err
	}

	delete(updateDoc, "_id")
	update := bson.D{{"$set", updateDoc}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updatedRule models.LoadBalancer
	err = l.getLoadBalancersCollection().FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedRule)
	if err != nil {
		return models.LoadBalancer{}, err
	}

	return updatedRule, nil
}
