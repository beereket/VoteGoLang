package domain

type Session struct {
	ID     uint
	UserID uint
	Expiry int64
}
