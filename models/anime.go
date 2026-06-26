package models

import (
	"time"
)

type Anime struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PosterURL   string    `json:"poseter_url"`
	BannerURL   string    `json:"banner_url"`
	Status      string    `json:"status"`
	ReleaseYear int       `json:"release_year"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
