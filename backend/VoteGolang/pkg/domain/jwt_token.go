package domain

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtToken struct {
	Secret []byte
}

type JwtClaims struct {
	SessionID uint `json:"sid"`
	UserID    uint `json:"uid"`
	jwt.StandardClaims
}

func NewJwtToken(secret string) *JwtToken {
	return &JwtToken{Secret: []byte(secret)}
}

func (tk *JwtToken) Create(s *Session, exp int64) (string, error) {
	claims := JwtClaims{
		SessionID: s.ID,
		UserID:    s.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tk.Secret)
}

func (tk *JwtToken) Check(ctx context.Context, inputToken string) (bool, error) {
	payload := &JwtClaims{}
	_, err := jwt.ParseWithClaims(inputToken, payload, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return tk.Secret, nil
	})
	if err != nil {
		return false, err
	}

	if payload.SessionID == 0 || payload.UserID == 0 {
		return false, fmt.Errorf("invalid token claims")
	}

	return true, nil
}
