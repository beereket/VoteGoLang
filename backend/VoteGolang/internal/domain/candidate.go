package domain

type Candidate struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Party      string `json:"party"`
	Region     string `json:"region"`
	Age        int    `json:"age"`
	Type       string `json:"type"`
	Votes      int    `json:"votes"`
	Photo      string `json:"photo,omitempty"`
	Education  string `json:"education,omitempty"`
	ElectionID int    `json:"election_id"`
}

type VoteRequest struct {
	CandidateID   int    `json:"candidate_id"`
	CandidateType string `json:"candidate_type"`
}
