package main

import (
	"fmt"
	"log"

	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/service"
)

func main() {

	// establish new database connection
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	//close the connection once the main function goes out of scope
	defer db.Close()

	// Checking if the connection established or not by Pining
	err = db.Ping()
	if err != nil {
		log.Fatal("Error in database ping", err)
	}

	fmt.Println("established connection successfully !!")

	err = db.TestQuery()
	if err != nil {
		log.Fatal("error in running the query", err)
	}

	// For testing User insertion purpose
	/*
		user := models.User{
			Username:     "abhayrajbhatn",
			Email:        "bhatabhayraj1608@gmail.com",
			PasswordHash: "test123",
		}

		user2 := models.User{
			Username:     "raja@123",
			Email:        "raja123@gmail.com",
			PasswordHash: "test234",
		}
		user3 := models.User{
			Username:     "abhi@2010",
			Email:        "abhi23@gmail.com",
			PasswordHash: "Pass@123",
		}

		err = service.Registeruser(user, db)

		if err != nil {
			log.Println("Restration error", err)
		} else {
			log.Println("Restration was successful!!")
		}

		err = service.Registeruser(user2, db)

		if err != nil {
			log.Println("Restration error", err)
		} else {
			log.Println("Restration was successful!!")
		}

		err = service.Registeruser(user3, db)

		if err != nil {
			log.Println("Restration error", err)
		} else {
			log.Println("Restration was successful!!")
		}
	*/

	// Test the Login function

	login_user := models.User{
		Username:     "abhi@2010",
		Email:        "abhi23@gmail.com",
		PasswordHash: "pass@123",
	}

	err = service.Loginuser(login_user, db)

	if err != nil {
		log.Println("failed to login", err)
	} else {
		fmt.Println("Login was successful")
	}
}
