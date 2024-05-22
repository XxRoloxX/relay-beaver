package proxyevents

import "backend/internal/request"

type ProxyEventsService struct {
	Repo request.RequestRepository
}

func (service ProxyEventsService) HandleProxiedRequest(message EventMessage) error {
	proxiedRequest := message.ProxiedRequest
	err := service.Repo.Create(proxiedRequest)
	return err
}
