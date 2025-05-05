package main

import (
	"VoteGolang/internal/background"
	"VoteGolang/internal/handlers"
	"VoteGolang/internal/middleware"
	"VoteGolang/internal/service"
	"log"
	"net/http"

	"VoteGolang/internal/database"
	"VoteGolang/internal/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.InitDB()

	// Services and Handlers
	logService := service.NewActivityLogService()
	logHandler := handlers.NewActivityLogHandler(logService)

	adminService := service.NewAdminService()
	adminHandler := handlers.NewAdminHandler(adminService)

	analyticsService := service.NewAnalyticsService()
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService)

	candidateService := service.NewCandidateService(logService)
	candidateHandler := handlers.NewCandidateHandler(candidateService)

	electionService := service.NewElectionService()
	electionHandler := handlers.NewElectionHandler(electionService)

	userService := service.NewUserService(logService)
	userHandler := handlers.NewUserHandler(userService)

	petitionService := service.NewPetitionService()
	petitionHandler := handlers.NewPetitionHandler(petitionService)

	newsService := service.NewNewsService()
	newsHandler := handlers.NewNewsHandler(newsService)

	deps := routes.RouterDependencies{
		AdminHandler:     adminHandler,
		AnalyticsHandler: analyticsHandler,
		CandidateHandler: candidateHandler,
		LogHandler:       logHandler,
		UserHandler:      userHandler,
		ElectionHandler:  electionHandler,
		PetitionHandler:  petitionHandler,
		NewsHandler:      newsHandler,
	}

	r := routes.RegisterRoutes(deps)

	handler := middleware.CORSMiddleware(r)
	background.StartElectionAutoCloser()

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
