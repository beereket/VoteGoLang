package domain

import "time"

type News struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Paragraph string    `json:"paragraph"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateNewsInput struct {
	Title     string `json:"title"`
	Paragraph string `json:"paragraph"`
	Photo     string `json:"photo"`
}
