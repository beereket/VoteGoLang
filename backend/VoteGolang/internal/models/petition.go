package models

type Petition struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
	VoteInFavor int    `json:"vote_in_favor"`
	VoteAgainst int    `json:"vote_against"`
	CreatedAt   string `json:"created_at"`
}
