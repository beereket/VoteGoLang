package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"database/sql"
	"time"
)

type AdminService struct{}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) GetDashboard() (*domain.DashboardData, error) {
	data := &domain.DashboardData{}

	err := database.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE deleted_at IS NULL`).Scan(&data.TotalUsers)
	if err != nil {
		return nil, err
	}
	err = database.DB.QueryRow(`SELECT COUNT(*) FROM votes WHERE deleted_at IS NULL`).Scan(&data.TotalVotes)
	if err != nil {
		return nil, err
	}
	err = database.DB.QueryRow(`SELECT COUNT(*) FROM candidates WHERE deleted_at IS NULL`).Scan(&data.TotalCandidates)
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(`
		SELECT name, votes FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY votes DESC LIMIT 3
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.TopCandidate
		if err := rows.Scan(&c.Name, &c.Votes); err != nil {
			return nil, err
		}
		data.TopCandidates = append(data.TopCandidates, c)
	}

	return data, nil
}

func (s *AdminService) GetVotesPerDay() ([]domain.VoteDay, error) {
	rows, err := database.DB.Query(`
		SELECT DATE(created_at), COUNT(*) FROM votes
		WHERE deleted_at IS NULL
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at) DESC LIMIT 7
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.VoteDay
	for rows.Next() {
		var v domain.VoteDay
		if err := rows.Scan(&v.Date, &v.Count); err != nil {
			return nil, err
		}
		result = append(result, v)
	}
	return result, nil
}

func (s *AdminService) GetUserRegistrationsPerWeek() ([]domain.UserWeek, error) {
	rows, err := database.DB.Query(`
		SELECT YEARWEEK(created_at, 1), COUNT(*) FROM users
		WHERE deleted_at IS NULL
		GROUP BY YEARWEEK(created_at, 1)
		ORDER BY YEARWEEK(created_at, 1) DESC LIMIT 4
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.UserWeek
	for rows.Next() {
		var u domain.UserWeek
		if err := rows.Scan(&u.YearWeek, &u.Count); err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func (s *AdminService) BanUser(userID int) error {
	_, err := database.DB.Exec(`UPDATE users SET deleted_at = ? WHERE id = ? AND deleted_at IS NULL`, time.Now(), userID)
	return err
}

func (s *AdminService) UnbanUser(userID int) error {
	result, err := database.DB.Exec(`UPDATE users SET deleted_at = NULL, updated_at = NOW() WHERE id = ? AND deleted_at IS NOT NULL`, userID)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (s *AdminService) ListAllUsers() ([]domain.User, error) {
	rows, err := database.DB.Query(`SELECT id, username, user_full_name, birth_date, address, deleted_at FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.User
	for rows.Next() {
		var u domain.User
		var deletedAt sql.NullTime
		if err := rows.Scan(&u.ID, &u.Username, &u.UserFullName, &u.BirthDate, &u.Address, &deletedAt); err != nil {
			return nil, err
		}
		if deletedAt.Valid {
			t := deletedAt.Time.Format(time.RFC3339)
			u.DeletedAt = &t
		}
		result = append(result, u)
	}
	return result, nil
}
