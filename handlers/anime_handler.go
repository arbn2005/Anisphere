package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/service"
)

func RegisterAnimeHandler(db *database.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Allow only POST requests.
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON request body.
		var anime models.Anime

		err := json.NewDecoder(r.Body).Decode(&anime)
		if err != nil {
			http.Error(w, "Invalid JSON Body", http.StatusBadRequest)
			return
		}

		// Call the service layer.
		err = service.Registeranime(anime, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Send success response.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Anime added successfully",
		})
	}
}
