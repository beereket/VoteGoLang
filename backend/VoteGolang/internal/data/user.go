package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username     string `json:"username"`
	UserFullName string `json:"user_full_name"`
	Password     string `json:"password"`
	BirthDate    string `json:"birth_date"`
	Address      string `json:"address"`
}
