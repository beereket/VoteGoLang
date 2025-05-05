package domain

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	UserFullName string `json:"user_full_name"`
	Password     string `json:"password"`
	BirthDate    string `json:"birth_date"`
	Address      string `json:"address"`
	Role         string `json:"role"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
