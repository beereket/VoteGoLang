package routes

import (
	"VoteGolang/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterRoutes(dep RouterDependencies) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ VoteGolang backend running"))
	})
	r.HandleFunc("/news", dep.NewsHandler.List).Methods("GET")

	r.HandleFunc("/users/create", dep.UserHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/login", dep.UserHandler.LoginUser).Methods("POST")

	r.HandleFunc("/elections/{election_id}/candidates", dep.CandidateHandler.ListByElection).Methods("GET")
	r.HandleFunc("/petitions", dep.PetitionHandler.List).Methods("GET")
	r.HandleFunc("/petitions/{id}", dep.PetitionHandler.GetByID).Methods("GET")
	r.HandleFunc("/petitions/{id}/results", dep.PetitionHandler.GetResults).Methods("GET")

	r.HandleFunc("/elections", dep.ElectionHandler.List).Methods("GET")
	r.HandleFunc("/elections/{id}", dep.ElectionHandler.Get).Methods("GET")
	r.HandleFunc("/elections/{id}/results", dep.ElectionHandler.GetResults).Methods("GET")

	r.Handle("/petitions/vote", middleware.UserAuthMiddleware(http.HandlerFunc(dep.PetitionHandler.Vote))).Methods("POST")
	r.Handle("/votes/cast", middleware.UserAuthMiddleware(http.HandlerFunc(dep.CandidateHandler.Vote))).Methods("POST")

	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AdminOnly)
	admin.HandleFunc("/logs", dep.LogHandler.List).Methods("GET")
	admin.HandleFunc("/users", dep.AdminHandler.ListAllUsers).Methods("GET")

	admin.HandleFunc("/dashboard", dep.AdminHandler.GetDashboard).Methods("GET")
	admin.HandleFunc("/dashboard/votes-per-day", dep.AdminHandler.GetVotesPerDay).Methods("GET")
	admin.HandleFunc("/dashboard/users-per-week", dep.AdminHandler.GetUserRegistrationsPerWeek).Methods("GET")
	admin.HandleFunc("/analytics/top-candidates", dep.AnalyticsHandler.GetTopCandidatesPerParty).Methods("GET")
	admin.HandleFunc("/analytics/party", dep.AnalyticsHandler.GetPartyAnalytics).Methods("GET")

	admin.HandleFunc("/elections", dep.ElectionHandler.List).Methods("GET")
	admin.HandleFunc("/elections/{election_id}/full", dep.ElectionHandler.GetWithCandidates).Methods("GET")
	admin.HandleFunc("/elections/create", dep.ElectionHandler.Create).Methods("POST")
	admin.HandleFunc("/elections/update/{id}", dep.ElectionHandler.Update).Methods("PUT")
	admin.HandleFunc("/elections/delete/{id}", dep.ElectionHandler.Delete).Methods("DELETE")

	admin.HandleFunc("/stats/users", dep.AnalyticsHandler.GetTotalUsers).Methods("GET")
	admin.HandleFunc("/stats/votes", dep.AnalyticsHandler.GetTotalVotes).Methods("GET")
	admin.HandleFunc("/stats/petitions", dep.AnalyticsHandler.GetTotalPetitions).Methods("GET")
	admin.HandleFunc("/stats/candidates", dep.AnalyticsHandler.GetTotalCandidates).Methods("GET")
	admin.HandleFunc("/stats/party-votes", dep.AnalyticsHandler.GetPartyVotes).Methods("GET")

	admin.HandleFunc("/users/ban/{id}", dep.AdminHandler.BanUser).Methods("DELETE")
	admin.HandleFunc("/users/unban/{id}", dep.AdminHandler.UnbanUser).Methods("PUT")
	admin.HandleFunc("/users/admins", dep.UserHandler.ListAdminUsers).Methods("GET")
	admin.HandleFunc("/users/create-admin", dep.UserHandler.CreateAdminUser).Methods("POST")

	admin.HandleFunc("/petitions/create", dep.PetitionHandler.Create).Methods("POST")
	admin.HandleFunc("/petitions/delete/{id}", dep.PetitionHandler.Delete).Methods("DELETE")

	admin.HandleFunc("/candidates/create", dep.CandidateHandler.Create).Methods("POST")
	admin.HandleFunc("/candidates/update/{id}", dep.CandidateHandler.Update).Methods("PUT")
	admin.HandleFunc("/candidates/delete/{id}", dep.CandidateHandler.Delete).Methods("DELETE")

	admin.HandleFunc("/news/create", dep.NewsHandler.Create).Methods("POST")
	admin.HandleFunc("/news/delete/{id}", dep.NewsHandler.Delete).Methods("DELETE")

	return r
}
