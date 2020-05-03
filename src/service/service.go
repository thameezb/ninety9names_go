package service

import (
	"fmt"
	"math/rand"

	"github.com/thameezb/ninety9names/src/models"
)

//GetAll returns a full list of the 99 names
func GetAll() map[int]models.Name {
	return models.Names
}

//GetName returns name associated with i, if no arg is sent, a random name is returned
func GetName(i int) (models.Name, error) {
	if i == -1 {
		i = rand.Intn(98)
	}
	if i > 99 || i < 1 {
		return models.Name{}, fmt.Errorf("number must be between 1-99")
	}
	return models.Names[i-1], nil
}
