package domain

type TopCandidate struct {
	Name  string
	Votes int
}

type DashboardData struct {
	TotalUsers      int
	TotalVotes      int
	TotalCandidates int
	TopCandidates   []TopCandidate
}

type VoteDay struct {
	Date  string
	Count int
}

type UserWeek struct {
	YearWeek string
	Count    int
}
