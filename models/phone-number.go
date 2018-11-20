package models

type PhoneNumber struct {
	Number      string `json:"number"`
	ClientInput string `json:"client_input"`
	Country     string `json:"country"`
}
