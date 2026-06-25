package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhayrajgithub1234/anisphere/auth"
	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/handlers"
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

	// Create a port that listens for incomming requests

	http.Handle(
		"/me",
		auth.AuthMiddleware(
			handlers.MeHandler(db),
		),
	)

	http.HandleFunc("/register", handlers.ResgisterHandler(db))
	http.HandleFunc("/login", handlers.LoginHandler(db))
	http.ListenAndServe(":8080", nil)

}
