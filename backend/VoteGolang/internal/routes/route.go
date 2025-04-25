package routes

import (
	"VoteGolang/internal/handlers"
	"database/sql"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
	return router
}
