package data

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Education string `json:"education"`
	Age       int    `json:"age"`
	Party     string `json:"party"`
	Region    string `json:"region"`
	Votes     int    `json:"votes" gorm:"default:0"`
	Type      string `json:"candidate_type"`
}
