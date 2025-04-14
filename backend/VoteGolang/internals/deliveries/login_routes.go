package deliveries

import (
	"VoteGolang/internals/services/auth"
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func LoginRegisterRoutes(mux *http.ServeMux, authHandler *auth.AuthHandler, tokenManager domain.TokenManager) {

	logRequest := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	//login
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		logLoginRegister(w, r, "/login")
		authHandler.Login(w, r)
	})

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		logLoginRegister(w, r, "/register")
		authHandler.Register(w, r)
	})

	//default
	mux.Handle("/", logRequest(mux))
}

func logLoginRegister(w http.ResponseWriter, r *http.Request, route string) {
	log.Printf("Attempting to access %s route, Method: %s, URL: %s", route, r.Method, r.URL.Path)
}
