package utils

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const usernameKey contextKey = "username"

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenStr := authHeader[len("Bearer "):]
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), usernameKey, claims.Username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func VerifyCredentials(db *sql.DB, username, password string) error {
	var storedHash string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username or password")
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		return errors.New("invalid username or password")
	}

	return nil
}
