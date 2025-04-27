package models

type Candidate struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Photo     string  `json:"photo"`
	Education string  `json:"education"`
	Age       int     `json:"age"`
	Party     string  `json:"party"`
	Region    string  `json:"region"`
	Votes     int     `json:"votes"`
	Type      string  `json:"type"`
	DeletedAt *string `json:"deleted_at,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
}
