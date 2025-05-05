package routes

import (
	"VoteGolang/internal/handlers"
)

type RouterDependencies struct {
	AdminHandler     *handlers.AdminHandler
	AnalyticsHandler *handlers.AnalyticsHandler
	CandidateHandler *handlers.CandidateHandler
	ElectionHandler  *handlers.ElectionHandler
	LogHandler       *handlers.ActivityLogHandler
	NewsHandler      *handlers.NewsHandler
	PetitionHandler  *handlers.PetitionHandler
	UserHandler      *handlers.UserHandler
}
