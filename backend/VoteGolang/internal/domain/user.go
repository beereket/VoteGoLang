package domain

import "time"

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	UserFullName string    `json:"user_full_name"`
	Password     string    `json:"password,omitempty"`
	BirthDate    string    `json:"birth_date"`
	Address      string    `json:"address"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    *string   `json:"deleted_at,omitempty"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
