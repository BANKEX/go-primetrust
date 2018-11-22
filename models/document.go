package models

type Document struct {
	AccountID   string `json:"account-id"`
	ContactID   string `json:"contact-id"`
	Description string `json:"description"`
	Extension   string `json:"extension"`
	File        string `json:"file"`
	Label       string `json:"label"`
	MimeType    string `json:"mime_type"`
	Public      bool   `json:"public"`
	//response only
	CreatedAt   string `json:"created-at"`
	FileURL     string `json:"file-url"`
	VersionUrls struct {
		Original string `json:"original"`
	} `json:"version-urls"`
}

type DocumentResponse struct {
	CollectionResponse
	Data DocumentData `json:"data"`
}

type DocumentData struct {
	Attributes    Document      `json:"attributes"`
	Type          string        `json:"type"`
	Links         Links         `json:"links"`
	Relationships Relationships `json:"relationships"`
	ID            string        `json:"id,omitempty"`
}

func NewDocument(accountId string, contactId string, file string, description string, extension string, label string, mimeType string, public bool) *Document {
	document := Document{
		AccountID:   accountId,
		ContactID:   contactId,
		Description: description,
		Extension:   extension,
		File:        file,
		Label:       label,
		MimeType:    mimeType,
		Public:      false,
	}
	return &document
}
