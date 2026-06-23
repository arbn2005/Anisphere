package main

import (
	"fmt"
	"log"

	"github.com/Abhayrajgithub1234/anisphere/database"
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

	fmt.Print("established connection successfully !!")

	err = db.TestQuery()
	if err != nil {
		log.Fatal("error in running the query", err)
	}
}
