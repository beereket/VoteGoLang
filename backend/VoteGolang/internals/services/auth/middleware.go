package auth

import (
	"VoteGolang/pkg/domain"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
)

type contextKey string

const sessionKey contextKey = "session"

func JWTMiddleware(tokenManager domain.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := ExtractTokenFromRequest(r)
			if err != nil {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			log.Printf("Token extracted: %s", token)

			valid, err := tokenManager.Check(r.Context(), token)
			if err != nil || !valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func ExtractTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid Authorization header format")
	}

	return parts[1], nil
}

func ExtractUserIDFromToken(token string, secret []byte) (string, error) {
	if token == "" {
		return "", errors.New("empty token")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", errors.New("user ID not found in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}
