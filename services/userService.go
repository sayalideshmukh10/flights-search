package services

import (
	"flights-server/models"
)

func ValidateCredentials(l models.Login) (bool, error) {

	if l.Username == "Admin" && l.Password == "Admin@123" {
		return true, nil
	}

	return false, nil

}
