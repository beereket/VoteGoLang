package handlers

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type User struct {
	Username     string `json:"username"`
	UserFullName string `json:"user_full_name"`
	Password     string `json:"password"`
	BirthDate    string `json:"birth_date"`
	Address      string `json:"address"`
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

		query := `
			INSERT INTO users (username, user_full_name, password, birth_date, address, created_at, updated_at)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`
		_, err := db.Exec(query, u.Username, u.UserFullName, string(hashedPassword), u.BirthDate, u.Address, time.Now(), time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}
