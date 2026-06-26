package service

import (
	"errors"

	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/repository"
)

func Registeranime(anime models.Anime, db *database.Database) error {
	// Validate
	if anime.Title == "" {
		return errors.New("title cannot be empty")
	}

	// Check duplicate
	existingAnime, err := repository.GetAnimeByTitle(anime.Title, db)
	if err != nil {
		return err
	}

	if existingAnime != nil {
		return errors.New("anime already exists")
	}

	// Register the new anime
	return repository.CreateAnime(anime, db)
}
