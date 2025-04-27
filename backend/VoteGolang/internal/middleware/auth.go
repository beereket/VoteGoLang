package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "❌ Authorization header missing", http.StatusUnauthorized)
			return
		}

		fmt.Println("🔍 Received Authorization Header:", authHeader)
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		fmt.Println("🔍 Parsed Token String:", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "❌ Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["role"] != "admin" {
			http.Error(w, "❌ Admin access only", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "username", claims["username"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
