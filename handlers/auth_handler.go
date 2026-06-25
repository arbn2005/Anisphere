package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Abhayrajgithub1234/anisphere/auth"
	"github.com/Abhayrajgithub1234/anisphere/database"
	"github.com/Abhayrajgithub1234/anisphere/models"
	"github.com/Abhayrajgithub1234/anisphere/service"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"Email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ResgisterHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqUser RegisterRequest

		err := json.NewDecoder(r.Body).Decode(&reqUser)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		user := models.User{
			Username:     reqUser.Username,
			Email:        reqUser.Email,
			PasswordHash: reqUser.Password,
		}

		err = service.Registeruser(user, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "user registed successfully",
			},
		)

	}
}

func LoginHandler(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginUser LoginRequest

		err := json.NewDecoder(r.Body).Decode(&loginUser)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		user := models.User{
			Email:        loginUser.Email,
			PasswordHash: loginUser.Password,
		}

		var logged_in_user *models.User

		logged_in_user, err = service.Loginuser(user, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var token string

		token, err = auth.GenerateToken(logged_in_user.ID)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
			"token":   token,
		})

	}
}
