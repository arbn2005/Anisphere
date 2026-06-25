package auth

import (
	"context"
	"net/http"
	"strings"
)

// Custom type to avoid key collisions in context.
type contextKey string

const UserIDKey contextKey = "userID"

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Read the Authorization header.
		authHeader := r.Header.Get("Authorization")

		// Ensure it begins with "Bearer ".
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Extract only the JWT.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT.
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store the user ID in the request context.
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)

		// Pass the modified request to the next handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
