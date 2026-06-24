package service

import (
	"errors"

	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/repository"
)

func Registeruser(user models.User, db *database.Database) error {

	_, err := repository.GetUserByEmail(user.Email, db)
	// If user exists return error - user already exists
	if err == nil {
		return errors.New("users already exists")
	}
	// If the user doesn't exists then create a new one by dereferencing the value
	err = repository.CreateUser(user, db)

	return err

}
