package service

import "ninety9names/models"

//GetAll returns a full list of the 99 names
func GetAll() map[int]models.Name {
	return models.Names
}

func GetName(i int) (models.Name, error) {

}
