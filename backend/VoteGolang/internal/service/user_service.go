package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"database/sql"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type UserService struct {
	LogService *ActivityLogService
}

func NewUserService(logService *ActivityLogService) *UserService {
	return &UserService{LogService: logService}
}

func (s *UserService) Register(user domain.User, isAdmin bool) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return 0, err
	}

	role := "guest"
	if isAdmin {
		role = "admin"
	}

	query := `
	INSERT INTO users (username, user_full_name, password, birth_date, address, role, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	r, err := database.DB.Exec(query,
		user.Username, user.UserFullName, hashedPassword,
		user.BirthDate, user.Address, role,
		time.Now(), time.Now(),
	)
	if err != nil {
		return 0, err
	}
	id, _ := r.LastInsertId()
	s.LogService.Create(int(id), "registration", fmt.Sprintf("%s registered", role))
	return int(id), nil
}

func (s *UserService) Login(input domain.LoginInput) (string, string, error) {
	var id int
	var hashedPassword, role string
	var deletedAt sql.NullTime

	err := database.DB.QueryRow(`SELECT id, password, role, deleted_at FROM users WHERE username = ?`, input.Username).
		Scan(&id, &hashedPassword, &role, &deletedAt)
	if err != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}
	if deletedAt.Valid {
		return "", "", fmt.Errorf("account is banned")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(input.Password)) != nil {
		return "", "", fmt.Errorf("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id":  id,
		"username": input.Username,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}
	s.LogService.Create(id, "login", fmt.Sprintf("User '%s' logged in", input.Username))
	return token, role, nil
}

func (s *UserService) ListAdmins() ([]domain.User, error) {
	rows, err := database.DB.Query(`SELECT id, username, user_full_name, birth_date, address FROM users WHERE role = 'admin' AND deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		if err := rows.Scan(&u.ID, &u.Username, &u.UserFullName, &u.BirthDate, &u.Address); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
