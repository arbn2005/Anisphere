package repository

import (
	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
)

func CreateUser(User models.User, db *database.Database) error {
	query := `
						INSERT INTO users (
							username,
							email,
							password_hash
						)
						VALUES ($1, $2, $3)
	`

	_, err := db.Conn.Exec(
		query,
		User.Username,
		User.Email,
		User.PasswordHash,
	)
	return err
}

func GetUserByEmail(email string, db *database.Database) (*models.User, error) {
	query := `
					SELECT id, username, email, password_hash
					FROM users
					WHERE users.email = $1
	`
	var user models.User
	err := db.Conn.QueryRow(
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	} else {
		return &user, err
	}
}
