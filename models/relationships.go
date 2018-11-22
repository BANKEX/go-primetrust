package models

type Relationship struct {
	Links Links `json:"links"`
}

type Relationships struct {
	Accounts                 Relationship `json:"accounts,omitempty"`
	Account                  Relationship `json:"account,omitempty"`
	Contacts                 Relationship `json:"contacts,omitempty"`
	Contact                  Relationship `json:"contact,omitempty"`
	AccountType              Relationship `json:"account-type,omitempty"`
	Addresses                Relationship `json:"addresses,omitempty"`
	AMLChecks                Relationship `json:"aml-checks,omitempty"`
	CIPChecks                Relationship `json:"cip-checks,omitempty"`
	Contributions            Relationship `json:"contributions,omitempty"`
	Disbursements            Relationship `json:"disbursements,omitempty"`
	FromContactRelationships Relationship `json:"from-contact-relationships,omitempty"`
	PaymentMethods           Relationship `json:"payment-methods,omitempty"`
	PhoneNumbers             Relationship `json:"phone-numbers,omitempty"`
	UploadedDocuments        Relationship `json:"uploaded-documents,omitempty"`
	RelatedFromContacts      Relationship `json:"related-from-contacts,omitempty"`
	RelatedToContacts        Relationship `json:"related-to-contacts,omitempty"`
	ToContactRelationships   Relationship `json:"to-contact-relationships,omitempty"`
	PrimaryAddress           Relationship `json:"primary-address,omitempty"`
	PrimaryContact           Relationship `json:"primary-contact,omitempty"`
	PrimaryPhoneNumber       Relationship `json:"primary-phone-number,omitempty"`
}
