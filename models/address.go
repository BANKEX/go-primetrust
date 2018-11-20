package models

type Address struct {
	city        string `json:"city"`
	country     string `json:"country"`
	postal_code string `json:"postal_code"`
	region      string `json:"region"`
	street_1    string `json:"street_1"`
	street_2    string `json:"street_2,omitempty"`
}
