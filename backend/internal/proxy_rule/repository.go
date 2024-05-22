package proxyrule

import (
	"backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ProxyRuleCollection = "proxy_rules"
)

type ProxyRuleMongoRepository struct {
	Db mongo.Database
}

func (r ProxyRuleMongoRepository) getProxyRuleCollection() *mongo.Collection {
	return r.Db.Collection(ProxyRuleCollection)
}

type ProxyRuleRepository interface {
	Create(proxyRule models.ProxyRule) (models.ProxyRule, error)
	FindById(id string) (models.ProxyRule, error)
	FindAll() ([]models.ProxyRule, error)
	Update(id string, proxyRule models.ProxyRule) (models.ProxyRule, error)
	Delete(id string) error
}

func (r ProxyRuleMongoRepository) Create(proxyRule models.ProxyRule) (models.ProxyRule, error) {
	res, error := r.getProxyRuleCollection().InsertOne(context.TODO(), proxyRule)

	if error != nil {
		return models.ProxyRule{}, error
	}

	proxyRule.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return proxyRule, nil
}

func (r ProxyRuleMongoRepository) FindById(id string) (models.ProxyRule, error) {

	primitiveId, idError := primitive.ObjectIDFromHex(id)

	if idError != nil {
		return models.ProxyRule{}, idError
	}

	res := r.getProxyRuleCollection().FindOne(context.TODO(), bson.M{"_id": primitiveId})

	if res.Err() != nil {
		return models.ProxyRule{}, res.Err()
	}

	var proxyRule models.ProxyRule

	err := res.Decode(&proxyRule)

	if err != nil {
		return models.ProxyRule{}, err
	}

	return proxyRule, nil
}

func (r ProxyRuleMongoRepository) FindAll() ([]models.ProxyRule, error) {
	entries, erro := r.getProxyRuleCollection().Find(context.TODO(), bson.D{})
	if erro != nil {
		return nil, erro
	}

	var proxyRuleEntries []models.ProxyRule

	for entries.Next(context.TODO()) {
		var proxyRuleEntry models.ProxyRule
		erro := entries.Decode(&proxyRuleEntry)
		if erro != nil {
			return nil, erro
		}
		proxyRuleEntries = append(proxyRuleEntries, proxyRuleEntry)

	}

	return proxyRuleEntries, nil
}

func (r ProxyRuleMongoRepository) Update(id string, proxyRuleEntry models.ProxyRule) (models.ProxyRule, error) {
	modifiedEntry := r.getProxyRuleCollection().FindOneAndUpdate(context.TODO(), id, proxyRuleEntry)

	if modifiedEntry.Err() != nil {
		return models.ProxyRule{}, modifiedEntry.Err()
	}
	var updatedEntry models.ProxyRule

	modifiedEntry.Decode(&updatedEntry)

	return updatedEntry, nil
}

func (r ProxyRuleMongoRepository) Delete(id string) error {
	primitiveId, idError := primitive.ObjectIDFromHex(id)

	if idError != nil {
		return idError
	}
	_, err := r.getProxyRuleCollection().DeleteOne(context.TODO(), bson.M{"_id": primitiveId})

	return err
}
