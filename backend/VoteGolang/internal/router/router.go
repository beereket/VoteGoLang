package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"VoteGolang/internal/handler"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Admin routes (no token for now)
	admin := r.Group("/admin")
	{
		admin.GET("/votes", handler.GetAllVotes(db))
		admin.GET("/users", handler.GetAllUsers(db))
		admin.POST("/petitions", handler.CreatePetition(db))
	}

	return r
}
