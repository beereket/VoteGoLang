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
	r.HandleFunc("/news", handlers.ListNews).Methods("GET")

	r.HandleFunc("/users/create", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/login", handlers.LoginUser).Methods("POST")

	r.HandleFunc("/elections/{election_id}/candidates", handlers.ListCandidatesByElection).Methods("GET")
	r.HandleFunc("/petitions", handlers.ListPetitions).Methods("GET")
	r.HandleFunc("/petitions/{id}", handlers.GetPetitionByID).Methods("GET")
	r.HandleFunc("/petitions/{id}/results", handlers.GetPetitionResults).Methods("GET")

	r.HandleFunc("/elections", handlers.ListElections).Methods("GET")
	r.HandleFunc("/elections/{id}", handlers.GetElectionDetail).Methods("GET")
	r.HandleFunc("/elections/{id}/results", handlers.GetElectionResults).Methods("GET")

	r.Handle("/petitions/vote", middleware.UserAuthMiddleware(http.HandlerFunc(handlers.VoteOnPetition))).Methods("POST")
	r.Handle("/votes/cast", middleware.UserAuthMiddleware(http.HandlerFunc(handlers.VoteForCandidate))).Methods("POST")

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AdminOnly)
	admin.HandleFunc("/users", handlers.ListAllUsers).Methods("GET")
	admin.HandleFunc("/logs", handlers.ListActivityLogs).Methods("GET")

	admin.HandleFunc("/dashboard", handlers.GetAdminDashboard).Methods("GET")
	admin.HandleFunc("/dashboard/votes-per-day", handlers.GetVotesPerDay).Methods("GET")
	admin.HandleFunc("/dashboard/users-per-week", handlers.GetUserRegistrationsPerWeek).Methods("GET")
	admin.HandleFunc("/analytics/top-candidates", handlers.GetTopCandidatesPerParty).Methods("GET")
	admin.HandleFunc("/analytics/election", handlers.GetElectionAnalytics).Methods("GET")
	admin.HandleFunc("/analytics/candidates", handlers.GetCandidateAnalytics).Methods("GET")
	admin.HandleFunc("/analytics/party", handlers.GetPartyAnalytics).Methods("GET")
	admin.HandleFunc("/analytics/party-percentage", handlers.GetPartyPercentageAnalytics).Methods("GET")

	admin.HandleFunc("/elections", handlers.ListElections).Methods("GET")
	admin.HandleFunc("/elections/{election_id}/full", handlers.GetElectionWithCandidates).Methods("GET")
	admin.HandleFunc("/elections/create", handlers.CreateElection).Methods("POST")
	admin.HandleFunc("/elections/update/{id}", handlers.UpdateElection).Methods("PUT")
	admin.HandleFunc("/elections/delete/{id}", handlers.DeleteElection).Methods("DELETE")
	admin.HandleFunc("/candidates/create", handlers.DeleteElection).Methods("POST")

	admin.HandleFunc("/stats/users", handlers.GetTotalUsers).Methods("GET")
	admin.HandleFunc("/stats/votes", handlers.GetTotalVotes).Methods("GET")
	admin.HandleFunc("/stats/petitions", handlers.GetTotalPetitions).Methods("GET")
	admin.HandleFunc("/stats/candidates", handlers.GetTotalCandidates).Methods("GET")
	admin.HandleFunc("/stats/party-votes", handlers.GetPartyVotes).Methods("GET")

	admin.HandleFunc("/users/ban/{id}", handlers.BanUser).Methods("DELETE")
	admin.HandleFunc("/users/unban/{id}", handlers.UnbanUser).Methods("PUT")
	admin.HandleFunc("/users/admins", handlers.ListAdminUsers).Methods("GET")
	admin.HandleFunc("/users/create-admin", handlers.CreateAdminUser).Methods("POST")

	admin.HandleFunc("/petitions/create", handlers.CreatePetition).Methods("POST")
	admin.HandleFunc("/petitions/close", handlers.ClosePetition).Methods("PUT")
	admin.HandleFunc("/petitions/delete/{id}", handlers.DeletePetition).Methods("DELETE")

	admin.HandleFunc("/candidates/create", handlers.CreateCandidate).Methods("POST")
	admin.HandleFunc("/candidates/update/{id}", handlers.UpdateCandidate).Methods("PUT")
	admin.HandleFunc("/candidates/delete/{id}", handlers.DeleteCandidate).Methods("DELETE")

	admin.HandleFunc("/news/create", handlers.CreateNews).Methods("POST")
	admin.HandleFunc("/news/delete/{id}", handlers.DeleteNews).Methods("DELETE")

	return r
}
