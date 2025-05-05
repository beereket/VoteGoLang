package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"time"
)

type ActivityLogService struct{}

func NewActivityLogService() *ActivityLogService {
	return &ActivityLogService{}
}

func (s *ActivityLogService) List() ([]domain.ActivityLog, error) {
	rows, err := database.DB.Query(`
		SELECT al.id, al.user_id, u.username, al.action, al.description, al.created_at
		FROM activity_logs al
		LEFT JOIN users u ON al.user_id = u.id
		ORDER BY al.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []domain.ActivityLog
	for rows.Next() {
		var l domain.ActivityLog
		if err := rows.Scan(&l.ID, &l.UserID, &l.Username, &l.Action, &l.Description, &l.CreatedAt); err != nil {
			return nil, err
		}
		logs = append(logs, l)
	}
	return logs, nil
}

func (s *ActivityLogService) Create(userID int, action, description string) error {
	_, err := database.DB.Exec(`
		INSERT INTO activity_logs (user_id, action, description, created_at)
		VALUES (?, ?, ?, ?)
	`, userID, action, description, time.Now())
	return err
}
