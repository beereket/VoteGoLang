package service

import (
	"VoteGolang/internal/database"
	"time"
)

// CreateActivityLog inserts a new activity into activity_logs table
func CreateActivityLog(userID int, action, description string) error {
	_, err := database.DB.Exec(`
		INSERT INTO activity_logs (user_id, action, description, created_at)
		VALUES (?, ?, ?, ?)
	`, userID, action, description, time.Now())
	return err
}
