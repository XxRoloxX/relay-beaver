package proxyrule

import "backend/pkg/models"

type ProxyRuleService struct {
	Repo ProxyRuleRepository
}

func (service ProxyRuleService) CreateRule(proxyRule models.ProxyRule) (models.ProxyRule, error) {
	return service.Repo.Create(proxyRule)
}

func (service ProxyRuleService) UpdateRule(proxyRule models.ProxyRule) (models.ProxyRule, error) {
	return service.Repo.Update(proxyRule)
}
