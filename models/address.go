package models

import (
	"github.com/satori/go.uuid"
)

const (
	AddressTypeHome = "home"
)

type Address struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Type       string    `json:"type,omitempty"`
	City       string    `json:"city"`
	Country    string    `json:"country"`
	PostalCode string    `json:"postal_code,omitempty"`
	Region     string    `json:"region,omitempty"`
	Street1    string    `json:"street_1"`
	Street2    string    `json:"street_2,omitempty"`
}
