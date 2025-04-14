package deliveries

import (
	"VoteGolang/internals/handlers"
	"VoteGolang/internals/services/auth"
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func RegisterPetitionRoutes(mux *http.ServeMux, handler *handlers.PetitionHandler, tokenManager domain.TokenManager) {
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

	mux.Handle("/vote/petition/create", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote/petition/create", handler.CreatePetition),
	))

	mux.Handle("/vote/petition/all", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote/petition/all", handler.GetAllPetitions),
	))

	mux.Handle("/vote/petition/get", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote/petition/get", handler.GetPetitionByID),
	))

	mux.Handle("/vote/petition/vote", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote/petition/vote", handler.Vote),
	))

	mux.Handle("/vote/petition/delete", auth.JWTMiddleware(tokenManager)(
		logRequest("/vote/petition/delete", handler.DeletePetition),
	))
}
