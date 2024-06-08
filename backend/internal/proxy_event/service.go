package proxyevent

import "backend/internal/request"
import "backend/pkg/models"

type ProxyEventsService struct {
	Repo request.RequestRepository
}

func (service ProxyEventsService) HandleProxiedRequest(proxiedRequest models.ProxiedRequest) error {
	// println(proxiedRequest.Target)
	_, err := service.Repo.Create(proxiedRequest)
	return err
}
