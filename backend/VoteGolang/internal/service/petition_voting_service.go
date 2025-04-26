package service

import (
	"fmt"
	"time"

	"VoteGolang/internal/database"
)

type PetitionVotingService struct{}

func (s *PetitionVotingService) CastPetitionVote(userID int, petitionID int, voteType string) error {
	// Check if already voted
	var exists int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM petition_votes
		WHERE user_id = ? AND petition_id = ? AND deleted_at IS NULL
	`, userID, petitionID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists > 0 {
		return fmt.Errorf("❌ You have already voted on this petition")
	}

	// Insert the petition vote
	_, err = database.DB.Exec(`
		INSERT INTO petition_votes (user_id, petition_id, vote_type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, userID, petitionID, voteType, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Update petition vote counts
	if voteType == "in_favor" {
		_, err = database.DB.Exec(`
			UPDATE petitions
			SET vote_in_favor = vote_in_favor + 1
			WHERE id = ?
		`, petitionID)
	} else if voteType == "against" {
		_, err = database.DB.Exec(`
			UPDATE petitions
			SET vote_against = vote_against + 1
			WHERE id = ?
		`, petitionID)
	}

	return err
}
