package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func Connection() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error in Conneting to DB", err)
	}
	// returns the struct embedded db poiner to the main function
	return &Database{
		Conn: db,
	}, nil
}

func (d *Database) Ping() error {
	err := d.Conn.Ping()
	return err
}

func (d *Database) Close() error {
	err := d.Conn.Close()
	return err
}

func (d *Database) TestQuery() error {
	var res int

	err := d.Conn.QueryRow("SELECT 1").Scan(&res)

	if err != nil {
		return err
	}

	fmt.Print("result: ", res)

	return nil
}
