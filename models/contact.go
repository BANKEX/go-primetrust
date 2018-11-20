package models

import (
	"github.com/satori/go.uuid"
)

const (
	ContactTypeNaturalPerson = "natural_person"
	ContactTypeCompany       = "company"
)

type ContactAttributeData struct {
	RelatedContacts []RelatedContactData `json:"related_contacts"`
}

type ContactAttributes struct {
	AccountRoles      []string             `json:"account-roles"`
	AMLCleared        bool                 `json:"aml-cleared"`
	CIPCleared        bool                 `json:"cip-cleared"`
	ContactType       string               `json:"contact-type"`
	Data              ContactAttributeData `json:"data"`
	DateOfBirth       string               `json:"date-of-birth"`
	Email             string               `json:"email"`
	Name              string               `json:"name"`
	RegionOfFormation string               `json:"region-of-formation"`
	Sex               string               `json:"sex"`
	TaxCountry        string               `json:"tax-country"`
	TaxIDNumber       string               `json:"tax-id-number"`
	TaxState          string               `json:"tax-state"`
}

type ContactData struct {
	Type          string            `json:"type"`
	ID            uuid.UUID         `json:"id"`
	Attributes    ContactAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type RelatedContactData struct {
	Type               string      `json:"type"`
	DateOfBirth        string      `json:"date_of_birth"`
	Email              string      `json:"email"`
	Name               string      `json:"name"`
	PrimaryAddress     Address     `json:"primary_address"`
	PrimaryPhoneNumber PhoneNumber `json:"primary_phone_number"`
	Sex                string      `json:"sex"`
	TaxCountry         string      `json:"tax_country"`
	TaxIDNumber        string      `json:"tax_id_number"`
	Label              string      `json:"label"`
}

type Contact struct {
	Data ContactData `json:"data"`
}

func NewNaturalPersonContact() *Contact {
	contact := Contact{
		Data: ContactData{
			Type: ContactTypeNaturalPerson,
		},
	}
	return &contact
}

func NewCompanyContact() *Contact {
	contact := Contact{
		Data: ContactData{
			Type: ContactTypeCompany,
		},
	}
	return &contact
}

type ContactsResponse struct {
	CollectionResponse
	Data []ContactData `json:"data"`
}
