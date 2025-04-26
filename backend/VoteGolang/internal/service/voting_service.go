package service

import (
	"fmt"
	"time"

	"VoteGolang/internal/database"
)

type VotingService struct{}

func (s *VotingService) CastVote(userID int, candidateID int, candidateType string) error {
	// 1. Check if already voted
	var exists int
	err := database.DB.QueryRow(`
		SELECT COUNT(*) FROM votes
		WHERE user_id = ? AND candidate_id = ? AND deleted_at IS NULL
	`, userID, candidateID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists > 0 {
		return fmt.Errorf("❌ You have already voted for this candidate")
	}

	// 2. Insert the vote
	_, err = database.DB.Exec(`
		INSERT INTO votes (user_id, candidate_id, candidate_type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, userID, candidateID, candidateType, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// 3. Update candidate vote count
	_, err = database.DB.Exec(`
		UPDATE candidates
		SET votes = votes + 1
		WHERE id = ?
	`, candidateID)
	if err != nil {
		return err
	}

	return nil
}
