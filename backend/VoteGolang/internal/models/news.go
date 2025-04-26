package models

type News struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Paragraph string `json:"paragraph"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
}
