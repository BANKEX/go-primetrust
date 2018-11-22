package models

import (
	"time"
)

type AccountAttributes struct {
	CreatedAt time.Time `json:"created-at"`
	Name      string    `json:"name"`
	Number    string    `json:"number"`
	Status    string    `json:"status"`
}

type AccountData struct {
	ID            string            `json:"id,omitempty"`
	Type          string            `json:"type"`
	Attributes    AccountAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type Account struct {
	Data AccountData `json:"data"`
}

type AccountsResponse struct {
	CollectionResponse
	Data []AccountData `json:"data"`
}
