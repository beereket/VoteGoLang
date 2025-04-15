package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"VoteGolang/internal/handler"
	"VoteGolang/internal/repository"
	"VoteGolang/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Repositories
	voteRepo := &repository.VoteRepo{}
	userRepo := &repository.UserRepo{}
	petitionRepo := &repository.PetitionRepo{}
	candidateRepo := &repository.CandidateRepo{}

	// Services
	voteService := service.NewVoteService(voteRepo)
	userService := service.NewUserService(userRepo)
	petitionService := service.NewPetitionService(petitionRepo)
	candidateService := service.NewCandidateService(candidateRepo)

	// Admin routes (no token for now)
	admin := r.Group("/admin")
	{
		admin.GET("/votes", handler.GetAllVotes(voteService))
		admin.GET("/users", handler.GetAllUsers(userService))
		admin.POST("/petitions", handler.CreatePetition(petitionService))
		admin.GET("/petitions", handler.GetAllPetitions(petitionService))
		admin.DELETE("/petitions/:id", handler.DeletePetition(petitionService))

		admin.GET("/candidates", handler.GetAllCandidates(candidateService))
		admin.DELETE("/votes/:id", handler.DeleteVote(voteService))
	}

	return r
}
