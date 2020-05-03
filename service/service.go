package service

import "ninety9names/models"

//ReturnAll returns a full list of the 99 names
func ReturnAll() []models.Name {
	return models.Names
}
