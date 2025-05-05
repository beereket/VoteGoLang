package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"time"
)

type ElectionService struct{}

func NewElectionService() *ElectionService {
	return &ElectionService{}
}

func (s *ElectionService) List() ([]domain.Election, error) {
	rows, err := database.DB.Query(`SELECT id, name, description, election_date FROM elections ORDER BY election_date ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var elections []domain.Election
	for rows.Next() {
		var e domain.Election
		if err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.ElectionDate); err != nil {
			return nil, err
		}
		elections = append(elections, e)
	}
	return elections, nil
}

func (s *ElectionService) Get(id int) (domain.Election, error) {
	var e domain.Election
	err := database.DB.QueryRow(`SELECT id, name, description, election_date FROM elections WHERE id = ?`, id).
		Scan(&e.ID, &e.Name, &e.Description, &e.ElectionDate)
	return e, err
}

func (s *ElectionService) Create(e domain.Election) error {
	_, err := database.DB.Exec(`
		INSERT INTO elections (name, description, election_date, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)`,
		e.Name, e.Description, e.ElectionDate, time.Now(), time.Now())
	return err
}

func (s *ElectionService) Update(id int, e domain.Election) error {
	_, err := database.DB.Exec(`
		UPDATE elections SET name = ?, description = ?, election_date = ?, updated_at = ? WHERE id = ?`,
		e.Name, e.Description, e.ElectionDate, time.Now(), id)
	return err
}

func (s *ElectionService) Delete(id int) error {
	_, err := database.DB.Exec(`DELETE FROM elections WHERE id = ?`, id)
	return err
}

func (s *ElectionService) GetWithCandidates(id int) (domain.ElectionWithCandidates, error) {
	var result domain.ElectionWithCandidates

	// Get election
	err := database.DB.QueryRow(`
		SELECT id, name, description, election_date
		FROM elections
		WHERE id = ?
	`, id).Scan(&result.Election.ID, &result.Election.Name, &result.Election.Description, &result.Election.ElectionDate)
	if err != nil {
		return result, err
	}

	// Get candidates
	rows, err := database.DB.Query(`
		SELECT id, name, party, region, age, type, votes, election_id
		FROM candidates
		WHERE election_id = ?
	`, id)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var c domain.Candidate
		if err := rows.Scan(&c.ID, &c.Name, &c.Party, &c.Region, &c.Age, &c.Type, &c.Votes, &c.ElectionID); err != nil {
			return result, err
		}
		result.Candidates = append(result.Candidates, c)
	}

	return result, nil
}

func (s *ElectionService) GetResults(electionID int) ([]domain.ElectionResult, error) {
	rows, err := database.DB.Query(`
		SELECT name, party, votes
		FROM candidates
		WHERE election_id = ?
		ORDER BY votes DESC
	`, electionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.ElectionResult
	for rows.Next() {
		var r domain.ElectionResult
		if err := rows.Scan(&r.Name, &r.Party, &r.Votes); err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, nil
}
