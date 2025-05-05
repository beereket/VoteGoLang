package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"time"
)

type NewsService struct{}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (s *NewsService) List(page, limit int) ([]domain.News, error) {
	offset := (page - 1) * limit
	rows, err := database.DB.Query(`
		SELECT id, title, paragraph, photo, created_at
		FROM general_news
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var newsList []domain.News
	for rows.Next() {
		var n domain.News
		if err := rows.Scan(&n.ID, &n.Title, &n.Paragraph, &n.Photo, &n.CreatedAt); err != nil {
			return nil, err
		}
		newsList = append(newsList, n)
	}

	return newsList, nil
}

func (s *NewsService) Create(input domain.CreateNewsInput) error {
	_, err := database.DB.Exec(`
		INSERT INTO general_news (title, paragraph, photo, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, input.Title, input.Paragraph, input.Photo, time.Now(), time.Now())
	return err
}

func (s *NewsService) Delete(id int) error {
	_, err := database.DB.Exec(`
		UPDATE general_news SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, id)
	return err
}
