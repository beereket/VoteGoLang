package domain

import "time"

type Petition struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Photo       string     `json:"photo"`
	VoteInFavor int        `json:"vote_in_favor"`
	VoteAgainst int        `json:"vote_against"`
	Status      string     `json:"status"`
	Deadline    *time.Time `json:"deadline,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type PetitionVoteRequest struct {
	PetitionID int    `json:"petition_id"`
	VoteType   string `json:"vote_type"`
}
