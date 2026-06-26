package repository

import (
	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
)

func CreateAnime(Anime models.Anime, db *database.Database) error {
	query := `
		INSERT INTO anime (
			title,
    description,
    poster_url,
    banner_url,
    status,
    release_year
		)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`

	_, err := db.Conn.Exec(
		query,
		Anime.Title,
		Anime.Description,
		Anime.PosterURL,
		Anime.BannerURL,
		Anime.Status,
		Anime.ReleaseYear,
	)

	return err
}

func GetAnimeByTitle(title string, db *database.Database) (*models.Anime, error) {
	query := `
		SELECT id, title, description, poster_url, banner_url, status, release_year, created_at, updated_at
		FROM anime
		WHERE anime.title = $1
	`
	var anime models.Anime
	err := db.Conn.QueryRow(query, title).Scan(
		&anime.ID,
		&anime.Title,
		&anime.Description,
		&anime.PosterURL,
		&anime.BannerURL,
		&anime.Status,
		&anime.ReleaseYear,
		&anime.CreatedAt,
		&anime.UpdatedAt,
	)
	if err != nil {
		return nil, err
	} else {
		return &anime, nil
	}
}
