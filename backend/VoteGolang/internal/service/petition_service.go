package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"database/sql"
	"errors"
)

type PetitionService struct{}

func NewPetitionService() *PetitionService {
	return &PetitionService{}
}

func (s *PetitionService) List() ([]domain.Petition, error) {
	rows, err := database.DB.Query(`SELECT id, user_id, title, description, photo, vote_in_favor, vote_against, status, deadline, created_at, updated_at FROM petitions ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var petitions []domain.Petition
	for rows.Next() {
		var p domain.Petition
		var deadline sql.NullTime
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.Photo, &p.VoteInFavor, &p.VoteAgainst, &p.Status, &deadline, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		if deadline.Valid {
			p.Deadline = &deadline.Time
		}
		petitions = append(petitions, p)
	}
	return petitions, nil
}

func (s *PetitionService) GetByID(id int) (domain.Petition, error) {
	var p domain.Petition
	var deadline sql.NullTime
	err := database.DB.QueryRow(`SELECT id, user_id, title, description, photo, vote_in_favor, vote_against, status, deadline, created_at, updated_at FROM petitions WHERE id = ?`, id).
		Scan(&p.ID, &p.UserID, &p.Title, &p.Description, &p.Photo, &p.VoteInFavor, &p.VoteAgainst, &p.Status, &deadline, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return p, err
	}
	if deadline.Valid {
		p.Deadline = &deadline.Time
	}
	return p, nil
}

func (s *PetitionService) Create(p domain.Petition) error {
	_, err := database.DB.Exec(`INSERT INTO petitions (user_id, title, description, photo, status, deadline, vote_in_favor, vote_against, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, 0, 0, NOW(), NOW())`, p.UserID, p.Title, p.Description, p.Photo, p.Status, p.Deadline)
	return err
}

func (s *PetitionService) Delete(id int) error {
	_, err := database.DB.Exec(`DELETE FROM petitions WHERE id = ?`, id)
	return err
}

func (s *PetitionService) Vote(userID int, req domain.PetitionVoteRequest) error {
	var exists int
	db := database.DB
	err := db.QueryRow(`SELECT COUNT(*) FROM petition_votes WHERE user_id = ? AND petition_id = ?`, userID, req.PetitionID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("⚠️ Already voted for this petition")
	}

	_, err = db.Exec(`INSERT INTO petition_votes (user_id, petition_id, vote_type, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`, userID, req.PetitionID, req.VoteType)
	if err != nil {
		return err
	}

	if req.VoteType == "favor" {
		_, err = db.Exec(`UPDATE petitions SET vote_in_favor = vote_in_favor + 1 WHERE id = ?`, req.PetitionID)
	} else if req.VoteType == "against" {
		_, err = db.Exec(`UPDATE petitions SET vote_against = vote_against + 1 WHERE id = ?`, req.PetitionID)
	} else {
		return errors.New("❌ Invalid vote type")
	}

	return err
}

func (s *PetitionService) GetResults(id int) (map[string]int, error) {
	var favor, against int
	err := database.DB.QueryRow(`SELECT vote_in_favor, vote_against FROM petitions WHERE id = ?`, id).Scan(&favor, &against)
	if err != nil {
		return nil, err
	}
	return map[string]int{"vote_in_favor": favor, "vote_against": against}, nil
}
