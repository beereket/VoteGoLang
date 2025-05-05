package domain

type CandidateAnalytics struct {
	CandidateName string
	CandidateType string
	Votes         int
}

type TopCandidateAnalytics struct {
	PartyName     string
	CandidateName string
	TotalVotes    int
}

type PartyAnalytics struct {
	PartyName  string
	TotalVotes int
	Percentage float64
}

type PartyVote struct {
	Party string
	Votes int
}
