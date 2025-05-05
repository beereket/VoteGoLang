package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"fmt"
	"log"
	"time"
)

type CandidateService struct {
	LogService *ActivityLogService
}

func NewCandidateService(logService *ActivityLogService) *CandidateService {
	return &CandidateService{
		LogService: logService,
	}
}

func (s *CandidateService) ListByElection(electionID string) ([]domain.Candidate, error) {
	rows, err := database.DB.Query(`SELECT id, name, party, region, age, type, votes FROM candidates WHERE election_id = ?`, electionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var candidates []domain.Candidate
	for rows.Next() {
		var c domain.Candidate
		if err := rows.Scan(&c.ID, &c.Name, &c.Party, &c.Region, &c.Age, &c.Type, &c.Votes); err != nil {
			return nil, err
		}
		candidates = append(candidates, c)
	}
	return candidates, nil
}

func (s *CandidateService) Create(c domain.Candidate) error {
	_, err := database.DB.Exec(`
		INSERT INTO candidates (name, photo, education, age, party, region, type, election_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, c.Name, c.Photo, c.Education, c.Age, c.Party, c.Region, c.Type, c.ElectionID, time.Now(), time.Now())
	return err
}

func (s *CandidateService) Update(id int, c domain.Candidate) error {
	_, err := database.DB.Exec(`
		UPDATE candidates SET name = ?, party = ?, region = ?, age = ?, type = ?, updated_at = NOW() WHERE id = ?
	`, c.Name, c.Party, c.Region, c.Age, c.Type, id)
	return err
}

func (s *CandidateService) Delete(id int) error {
	_, err := database.DB.Exec(`DELETE FROM candidates WHERE id = ?`, id)
	return err
}

func (s *CandidateService) CastVote(userID int, req domain.VoteRequest) error {
	var count int
	err := database.DB.QueryRow(`SELECT COUNT(*) FROM votes WHERE user_id = ? AND candidate_type = ?`, userID, req.CandidateType).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("already voted for this election")
	}

	_, err = database.DB.Exec(`
		INSERT INTO votes (user_id, candidate_id, candidate_type, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`, userID, req.CandidateID, req.CandidateType)
	if err != nil {
		return err
	}

	if err := s.LogService.Create(userID, "vote", fmt.Sprintf("Voted for candidate ID %d in %s", req.CandidateID, req.CandidateType)); err != nil {
		log.Println("⚠️ Failed to log vote activity:", err)
	}

	return nil
}
