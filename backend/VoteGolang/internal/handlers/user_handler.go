package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"

	"VoteGolang/internal/database"
	"VoteGolang/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, "❌ Failed to hash password", http.StatusInternalServerError)
		return
	}

	query := `
	INSERT INTO users (username, user_full_name, password, birth_date, address, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = database.DB.Exec(query,
		user.Username,
		user.UserFullName,
		hashedPassword,
		user.BirthDate,
		user.Address,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		http.Error(w, "❌ Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ User created"})
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var input LoginInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	var dbPassword, role string
	err := database.DB.QueryRow("SELECT password, role FROM users WHERE username = ?", input.Username).Scan(&dbPassword, &role)
	if err == sql.ErrNoRows {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "❌ Server error", http.StatusInternalServerError)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(input.Password)) != nil {
		http.Error(w, "❌ Incorrect password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": input.Username,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 24h expiry
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "❌ Failed to create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, "❌ Failed to hash password", http.StatusInternalServerError)
		return
	}

	query := `
	INSERT INTO users (username, user_full_name, password, birth_date, address, role, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = database.DB.Exec(query,
		user.Username,
		user.UserFullName,
		hashedPassword,
		user.BirthDate,
		user.Address,
		"admin",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		http.Error(w, "❌ Failed to create admin user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Admin user created"})
}

func ListAdminUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, username, user_full_name, birth_date, address
		FROM users
		WHERE role = 'admin' AND deleted_at IS NULL
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch admin users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var admins []models.User

	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.UserFullName, &u.BirthDate, &u.Address)
		if err != nil {
			http.Error(w, "❌ Error reading user", http.StatusInternalServerError)
			return
		}
		admins = append(admins, u)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(admins)
}
