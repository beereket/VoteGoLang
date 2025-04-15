package handler

import (
	"VoteGolang/internal/data"
	"VoteGolang/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllVotes(voteService service.VoteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		votes, err := voteService.GetAllVotes(c.MustGet("db").(*gorm.DB))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get votes"})
			return
		}
		c.JSON(http.StatusOK, votes)
	}
}

func GetAllUsers(userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userService.GetAllUsers(c.MustGet("db").(*gorm.DB))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func CreatePetition(petitionService service.PetitionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var petition data.Petition
		if err := c.ShouldBindJSON(&petition); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err := petitionService.CreatePetition(c.MustGet("db").(*gorm.DB), &petition); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create petition"})
			return
		}

		c.JSON(http.StatusCreated, petition)
	}
}
