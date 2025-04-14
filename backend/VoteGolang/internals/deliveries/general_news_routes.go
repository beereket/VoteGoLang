package deliveries

import (
	"VoteGolang/internals/handlers"
	"VoteGolang/internals/services/auth"
	"VoteGolang/pkg/domain"
	"log"
	"net/http"
)

func RegisterGeneralNewsRoutes(mux *http.ServeMux, handler *handlers.GeneralNewsHandler, tokenManager domain.TokenManager) {
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

	mux.Handle("/general_news", auth.JWTMiddleware(tokenManager)(
		logRequest("/general_news", handler.GetAll),
	))
	mux.Handle("/general_news/get", auth.JWTMiddleware(tokenManager)(
		logRequest("/general_news/get", handler.GetByID),
	))
	mux.Handle("/general_news/create", auth.JWTMiddleware(tokenManager)(
		logRequest("/general_news/create", handler.Create),
	))

	mux.Handle("/general_news/update", auth.JWTMiddleware(tokenManager)(
		logRequest("/general_news/update", handler.Update),
	))

	mux.Handle("/general_news/delete", auth.JWTMiddleware(tokenManager)(
		logRequest("/general_news/delete", handler.Delete),
	))
}
