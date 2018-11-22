package models

import (
	"github.com/satori/go.uuid"
)

type PhoneNumber struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Number      string    `json:"number"`
	SMS         bool      `json:"sms,omitempty"`
	ClientInput string    `json:"client_input,omitempty"`
	Country     string    `json:"country,omitempty"`
}
