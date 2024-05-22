package proxyrule

import (
	"backend/pkg/models"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProxyRuleMongoRepository struct {
	Db mongo.Database
}

type ProxyRuleRepository interface {
	Create(proxyRule models.ProxyRule) (models.ProxyRule, error)
	FindById(id string) (models.ProxyRule, error)
	FindAll() ([]models.ProxyRule, error)
	Update(proxyRule models.ProxyRule) (models.ProxyRule, error)
	Delete(id string) error
}

func (r ProxyRuleMongoRepository) Create(proxyRule models.ProxyRule) (models.ProxyRule, error) {
	res, error := r.Db.Collection("proxy_rules").InsertOne(context.TODO(), proxyRule)

	if error != nil {
		return models.ProxyRule{}, error
	}

	proxyRule.Id = res.InsertedID.(primitive.ObjectID).Hex()

	return proxyRule, nil
}

func (r ProxyRuleMongoRepository) FindById(id string) (models.ProxyRule, error) {
	return models.ProxyRule{}, nil
}

func (r ProxyRuleMongoRepository) FindAll() ([]models.ProxyRule, error) {
	entries, erro := r.Db.Collection("proxy_rules").Find(context.TODO(), nil)

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
	modifiedEntry := r.Db.Collection("proxy_rules").FindOneAndUpdate(context.TODO(), id, proxyRuleEntry)

	if modifiedEntry.Err() != nil {
		return models.ProxyRule{}, modifiedEntry.Err()
	}
	var updatedEntry models.ProxyRule

	modifiedEntry.Decode(&updatedEntry)

	return updatedEntry, nil
}

func (r ProxyRuleMongoRepository) Delete(id string) error {
	return nil
}
