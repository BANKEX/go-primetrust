package models

const (
	AddressTypeHome   = "home"
	AddressTypeOffice = "office"
)

type Address struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	City       string `json:"city"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code,omitempty"`
	Region     string `json:"region,omitempty"`
	Street1    string `json:"street_1"`
	Street2    string `json:"street_2,omitempty"`
}
