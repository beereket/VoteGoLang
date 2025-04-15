package data

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	UserID       uint   `json:"user_id" gorm:"not null"`
	Title        string `json:"title"`
	Photo        string `json:"photo"`
	Description  string `json:"description"`
	VotesInFavor int    `json:"vote_in_favor" gorm:"default:0"`
	VotesAgainst int    `json:"vote_against" gorm:"default:0"`
}
