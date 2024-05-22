package proxyrule

import "backend/pkg/models"

type ProxyRuleService struct {
	Repo ProxyRuleRepository
}

func (service ProxyRuleService) CreateRule(proxyRule models.ProxyRule) (models.ProxyRule, error) {
	return service.Repo.Create(proxyRule)
}

func (service ProxyRuleService) UpdateRule(id string, proxyRule models.ProxyRule) (models.ProxyRule, error) {
	return service.Repo.Update(id, proxyRule)
}

func (service ProxyRuleService) GetRules() ([]models.ProxyRule, error) {
	return service.Repo.FindAll()
}

func (service ProxyRuleService) DeleteRule(id string) error {
	return service.Repo.Delete(id)
}
func (service ProxyRuleService) isRulePresent(id string) bool {
	_, err := service.Repo.FindById(id)
	return err == nil
}
