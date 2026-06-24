package service

import (
	"errors"

	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/repository"
	"github.com/Abhayrajgithub1234/anisphere/utils"
)

func Registeruser(user models.User, db *database.Database) error {

	// The hashpassword passed present in the user is not hashed so it has to be hased to be saved
	// in the database
	hashpassword, err := utils.HashPassword(user.PasswordHash)

	if err != nil {
		return err
	}

	user.PasswordHash = hashpassword

	_, err = repository.GetUserByEmail(user.Email, db)
	// If user exists return error - user already exists
	if err == nil {
		return errors.New("users already exists")
	}
	// If the user doesn't exists then create a new one by dereferencing the value
	err = repository.CreateUser(user, db)

	return err

}

func Loginuser(user models.User, db *database.Database) error {

	var getUser *models.User

	getUser, err := repository.GetUserByEmail(user.Email, db)
	if err != nil {
		return err
	}

	res := utils.CheckPassword(getUser.PasswordHash, user.PasswordHash)

	if res == true {
		return nil
	} else {
		return errors.New("Wrong Password")
	}

}
