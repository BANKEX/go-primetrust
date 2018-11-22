package models

import (
	"github.com/satori/go.uuid"
)

const (
	ContactsType = "contacts"

	ContactTypeNaturalPerson = "natural_person"
	ContactTypeCompany       = "company"
	ContactType              = "c-corp"
	ContactTypeLLC           = "llc"
	ContactTypeSCorp         = "s-corp"
	ContactTypeTrust         = "trust"

	AccountRoleOwner            = "owner"
	AccountRoleTaxFormRecipient = "tax form recipient"
)

type ContactAttributeData struct {
	RelatedContacts []RelatedContactData `json:"related_contacts"`
}

type ContactAttributes struct {
	AccountID          uuid.UUID            `json:"account-id"`
	AccountRoles       []string             `json:"account-roles"`
	Type               string               `json:"type"`
	ContactType        string               `json:"contact-type,omitempty"`
	AMLCleared         bool                 `json:"aml-cleared"`
	CIPCleared         bool                 `json:"cip-cleared"`
	DateOfBirth        string               `json:"date-of-birth"`
	Email              string               `json:"email"`
	Name               string               `json:"name"`
	Sex                string               `json:"sex,omitempty"`
	Label              string               `json:"label,omitempty"`
	RegionOfFormation  string               `json:"region-of-formation,omitempty"`
	TaxCountry         string               `json:"tax-country"`
	TaxIDNumber        string               `json:"tax-id-number"`
	TaxState           string               `json:"tax-state,omitempty"`
	PrimaryAddress     Address              `json:"primary_address"`
	PrimaryPhoneNumber PhoneNumber          `json:"primary_phone_number"`
	Data               ContactAttributeData `json:"data"`
}

type ContactData struct {
	ID            uuid.UUID         `json:"id,omitempty"`
	Type          string            `json:"type"`
	Attributes    ContactAttributes `json:"attributes"`
	Links         Links             `json:"links"`
	Relationships Relationships     `json:"relationships"`
}

type RelatedContactData struct {
	Type               string      `json:"type"`
	DateOfBirth        string      `json:"date_of_birth"`
	Email              string      `json:"email"`
	Name               string      `json:"name"`
	Sex                string      `json:"sex"`
	Label              string      `json:"label"`
	TaxCountry         string      `json:"tax_country"`
	TaxIDNumber        string      `json:"tax_id_number"`
	PrimaryAddress     Address     `json:"primary_address"`
	PrimaryPhoneNumber PhoneNumber `json:"primary_phone_number"`
}

type Contact struct {
	Data ContactData `json:"data"`
}

func NewNaturalPersonContact(accountId uuid.UUID) *Contact {
	contact := Contact{
		Data: ContactData{
			Type: ContactsType,
			Attributes: ContactAttributes{
				AccountID:    accountId,
				AccountRoles: []string{AccountRoleOwner},
				Type:         ContactTypeNaturalPerson,
			},
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
