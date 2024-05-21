package proxyrule

import (
	"backend/pkg/models"
	"context"
	UUID "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProxyRuleMongoRepository struct {
	Db mongo.Database
}

type ProxyRuleRepository interface {
	Create(proxyRule models.ProxyRule) error
	FindById(id UUID.UUID) (models.ProxyRule, error)
	FindAll() ([]models.ProxyRule, error)
	Update(proxyRule models.ProxyRule) error
	Delete(id UUID.UUID) error
}

func (r ProxyRuleMongoRepository) Create(proxyRuleEntry models.ProxyRule) error {
	_, error := r.Db.Collection("proxy_rules").InsertOne(context.TODO(), proxyRuleEntry)

	return error
}

func (r ProxyRuleMongoRepository) FindById(id UUID.UUID) (models.ProxyRule, error) {
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

func (r ProxyRuleMongoRepository) Update(proxyRuleEntry models.ProxyRule) error {
	return nil
}

func (r ProxyRuleMongoRepository) Delete(id UUID.UUID) error {
	return nil
}
