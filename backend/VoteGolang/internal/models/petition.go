package models

import "time"

type Petition struct {
	ID           int        `json:"id"`
	UserID       int        `json:"user_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Photo        string     `json:"photo"`
	VotesInFavor int        `json:"votes_in_favor"`
	VotesAgainst int        `json:"votes_against"`
	Status       string     `json:"status"`             // <--- NEW
	Deadline     *time.Time `json:"deadline,omitempty"` // <--- NEW
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
