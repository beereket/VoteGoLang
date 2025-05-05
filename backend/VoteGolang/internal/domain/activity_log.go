package domain

type ActivityLog struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}
