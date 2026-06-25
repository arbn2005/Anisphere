package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abhayrajgithub1234/anisphere/auth"
	"github.com/Abhayrajgithub1234/anisphere/database"
)

func MeHandler(db *database.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		userID := r.Context().Value(auth.UserIDKey)

		json.NewEncoder(w).Encode(map[string]any{
			"message": "Authenticated",
			"userID":  userID,
		})
	}
}
