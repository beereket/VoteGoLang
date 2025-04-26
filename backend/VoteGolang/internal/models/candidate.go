package models

type Candidate struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Education string `json:"education"`
	Age       int    `json:"age"`
	Party     string `json:"party"`
	Region    string `json:"region"`
	Type      string `json:"type"`
}
