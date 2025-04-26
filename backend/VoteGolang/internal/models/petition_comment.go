package models

type PetitionComment struct {
	ID          int    `json:"id"`
	PetitionID  int    `json:"petition_id"`
	UserID      int    `json:"user_id"`
	CommentText string `json:"comment_text"`
	CreatedAt   string `json:"created_at"`
}
