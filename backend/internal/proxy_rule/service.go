package proxyrule

import "backend/pkg/models"

type ProxyRuleService struct {
	Repo ProxyRuleRepository
}

func (service ProxyRuleService) Create(proxyRule models.ProxyRule) error {
	return service.Repo.Create(proxyRule)
}
