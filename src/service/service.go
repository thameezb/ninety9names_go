package service

import (
	"errors"
	"math/rand"
	"strconv"

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

//GetName returns name associated with i, if r is sent, a random name is returned
func (s *Service) GetName(id string) (models.Name, error) {
	if id == "r" {
		id = strconv.Itoa(rand.Intn(98) + 1)
	} else {
		_, err := strconv.Atoi(id)
		if err != nil {
			return models.Name{}, errors.New("id must be between 1-99 or r for random")
		}
	}
	return s.DB.GetName(id)
}
