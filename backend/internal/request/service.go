package request

import (
	"backend/pkg/models"
)

type RequestService struct {
	Repo RequestRepository
}

func (service RequestService) CreateRequest(request models.Request) error {
	return service.Repo.Create(request)
}
func (service RequestService) FindById(id string) (models.Request, error) {
	return service.Repo.FindById(id)
}
func (service RequestService) FindAll() ([]models.Request, error) {
	return service.Repo.FindAll()
}
