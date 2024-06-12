package proxyevent

import "backend/internal/request"
import "backend/pkg/models"

type ProxyEventsService struct {
	Repo request.RequestRepository
}

func (service ProxyEventsService) HandleProxiedRequest(proxiedRequest models.ProxiedRequest) (models.ProxiedRequest, error) {
	newRequest, err := service.Repo.Create(proxiedRequest)
	if err != nil {
		return models.ProxiedRequest{}, err
	}
	return newRequest, nil
}
