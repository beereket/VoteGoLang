package domain

import "context"

type TokenManager interface {
	Create(*Session, int64) (string, error)
	Check(context.Context, string) (bool, error)
}
