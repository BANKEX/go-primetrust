package models

type PhoneNumber struct {
	ID          string `json:"id,omitempty"`
	Number      string `json:"number"`
	SMS         bool   `json:"sms,omitempty"`
	ClientInput string `json:"client-input,omitempty"`
	Country     string `json:"country,omitempty"`
}
