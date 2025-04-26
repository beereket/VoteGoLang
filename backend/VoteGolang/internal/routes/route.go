package routes

import (
	"VoteGolang/internal/handlers"
	"VoteGolang/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ VoteGolang backend running"))
	})

	r.HandleFunc("/users/create", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	r.HandleFunc("/candidates", handlers.ListCandidates).Methods("GET")

	r.HandleFunc("/results", handlers.ListElectionResults).Methods("GET")
	r.HandleFunc("/votes/cast", handlers.VoteForCandidate).Methods("POST")
	r.HandleFunc("/votes/history", handlers.GetUserVotingHistory).Methods("GET")

	r.HandleFunc("/petitions", handlers.ListPetitions).Methods("GET")
	r.HandleFunc("/petitions/create", handlers.CreatePetition).Methods("POST")
	r.HandleFunc("/petitions/vote", handlers.VoteOnPetition).Methods("POST")
	r.HandleFunc("/petitions/{id}/comments", handlers.ListPetitionComments).Methods("GET")
	r.HandleFunc("/petitions/comments/create", handlers.CreatePetitionComment).Methods("POST")

	// Protected admin routes:
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AdminOnly)
	admin.HandleFunc("/dashboard", handlers.GetAdminDashboard).Methods("GET")
	admin.HandleFunc("/dashboard/votes-per-day", handlers.GetVotesPerDay).Methods("GET")
	admin.HandleFunc("/dashboard/users-per-week", handlers.GetUserRegistrationsPerWeek).Methods("GET")

	admin.HandleFunc("/users/admins", handlers.ListAdminUsers).Methods("GET")
	admin.HandleFunc("/users/create-admin", handlers.CreateAdminUser).Methods("POST")

	admin.HandleFunc("/candidates/create", handlers.CreateCandidate).Methods("POST")
	admin.HandleFunc("/candidates/update/{id}", handlers.UpdateCandidate).Methods("PUT")
	admin.HandleFunc("/candidates/delete/{id}", handlers.DeleteCandidate).Methods("DELETE")

	admin.HandleFunc("/petitions/delete/{id}", handlers.DeletePetition).Methods("DELETE")
	admin.HandleFunc("/petitions/comments/delete/{id}", handlers.DeletePetitionComment).Methods("DELETE")
	//admin.HandleFunc("/news/create", handlers.CreateNews).Methods("POST")

	return r
}
