package deliveries

import (
	"VoteGolang/internals/handlers"
	"VoteGolang/internals/services/auth"
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func RegisterCandidateRoutes(mux *http.ServeMux, handler *handlers.CandidateHandler, tokenManager domain.TokenManager) {
	logRequest := func(route string, handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.ExtractTokenFromRequest(r)
			if err != nil {
				log.Printf("Failed to extract token: %v", err)
			} else {
				log.Printf("Accessing %s route | Method: %s | URL: %s | Token: %s", route, r.Method, r.URL.Path, token)
			}

			handlerFunc(w, r)
		}
	}

	mux.Handle("/candidates", auth.JWTMiddleware(tokenManager)(
		logRequest("/candidates", handler.GetAll),
	))

	mux.Handle("/vote", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote", handler.Vote),
	))
}
