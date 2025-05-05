package domain

type Election struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ElectionDate string `json:"election_date"`
}

type ElectionWithCandidates struct {
	Election   Election    `json:"election"`
	Candidates []Candidate `json:"candidates"`
}

type ElectionResult struct {
	Name  string `json:"name"`
	Party string `json:"party"`
	Votes int    `json:"votes"`
}
