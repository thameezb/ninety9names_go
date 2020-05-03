package service

import (
	"github.com/thameezb/ninety9names/src/models"
	"github.com/thameezb/ninety9names/src/repository"
)

type Interface interface {
	GetAll() ([]models.Name, error)
	GetName(id string) (models.Name, error)
}

type Service struct {
	DB repository.Interface
}

//GetAll returns a full list of the 99 names
func (s *Service) GetAll() ([]models.Name, error) {
	return s.DB.GetAll()
}

//GetName returns name associated with i, if no arg is sent, a random name is returned
func (s *Service) GetName(id string) (models.Name, error) {
	return s.DB.GetName(id)
}
