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
	conn *sql.DB
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
		log.Fatal("Error in conneting to DB", err)
	}
	// returns the struct embedded db poiner to the main function
	return &Database{
		conn: db,
	}, nil
}

func (d *Database) Ping() error {
	err := d.conn.Ping()
	return err
}

func (d *Database) Close() error {
	err := d.conn.Close()
	return err
}

func (d *Database) TestQuery() error {
	var res int

	err := d.conn.QueryRow("SELECT 1").Scan(&res)

	if err != nil {
		return err
	}

	fmt.Print("result: ", res)

	return nil
}
