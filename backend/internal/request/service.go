package request

import (
	"backend/pkg/models"
)

type RequestService struct {
	Repo RequestRepository
}

func (service RequestService) Store(request models.Request) error {
	return service.Repo.Store(request)
}
