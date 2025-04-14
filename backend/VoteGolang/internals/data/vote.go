package data

import "gorm.io/gorm"

type Vote struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	UserID        uint   `json:"user_id" gorm:"not null"`
	CandidateID   uint   `json:"candidate_id" gorm:"not null"`
	CandidateType string `json:"candidate_type"`
}
