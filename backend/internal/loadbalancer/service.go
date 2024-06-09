package loadbalancer

import "backend/pkg/models"

type LoadBalancerService struct {
	Repo LoadBalancerRepository
}

func (service LoadBalancerService) CreateLB(lb models.LoadBalancer) (models.LoadBalancer, error) {
	return service.Repo.Create(lb)
}

func (service LoadBalancerService) UpdateRule(id string, proxyRule models.LoadBalancer) (models.LoadBalancer, error) {
	return service.Repo.Update(id, proxyRule)
}

func (service LoadBalancerService) GetLBs() ([]models.LoadBalancer, error) {
	return service.Repo.FindAll()
}

func (service LoadBalancerService) DeleteLB(id string) error {
	return service.Repo.Delete(id)
}

func (service LoadBalancerService) isLBPresent(id string) bool {
	_, err := service.Repo.FindById(id)
	return err == nil
}
