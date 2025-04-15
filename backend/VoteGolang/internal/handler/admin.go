package handler

import (
	"VoteGolang/internal/data"
	"VoteGolang/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get all votes
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

// Get all users
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

// Create petition
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

// Get all petitions
func GetAllPetitions(petitionService service.PetitionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		petitions, err := petitionService.GetAllPetitions(c.MustGet("db").(*gorm.DB))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get petitions"})
			return
		}
		c.JSON(http.StatusOK, petitions)
	}
}

// Delete petition
func DeletePetition(petitionService service.PetitionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Convert string id to uint
		petitionID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid petition ID"})
			return
		}

		// Call the service method with the parsed uint ID
		if err := petitionService.DeletePetitionByID(c.MustGet("db").(*gorm.DB), uint(petitionID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete petition"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Petition deleted successfully"})
	}
}

// Get all candidates
func GetAllCandidates(candidateService service.CandidateService) gin.HandlerFunc {
	return func(c *gin.Context) {
		candidates, err := candidateService.GetAllCandidates(c.MustGet("db").(*gorm.DB))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get candidates"})
			return
		}
		c.JSON(http.StatusOK, candidates)
	}
}

// Delete vote
func DeleteVote(voteService service.VoteService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Convert string id to uint
		voteID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vote ID"})
			return
		}

		// Call the service method with the parsed uint ID
		if err := voteService.DeleteVoteByID(c.MustGet("db").(*gorm.DB), uint(voteID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vote"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Vote deleted successfully"})
	}
}
