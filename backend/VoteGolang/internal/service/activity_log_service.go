package service

import (
	"VoteGolang/internal/database"
	"time"
)

func CreateActivityLog(userID int, action, description string) error {
	_, err := database.DB.Exec(`
		INSERT INTO activity_logs (user_id, action, description, created_at)
		VALUES (?, ?, ?, ?)
	`, userID, action, description, time.Now())
	return err
}
