package models

import "mime/multipart"

type Document struct {
	AccountID   string         `json:"account-id,omitempty"`
	ContactID   string         `json:"contact-id,omitempty"`
	Description string         `json:"description,omitempty"`
	Extension   string         `json:"extension,omitempty"`
	File        multipart.File `json:"file,omitempty"`
	Label       string         `json:"label,omitempty"`
	MimeType    string         `json:"mime_type,omitempty"`
	Public      bool           `json:"public,omitempty"`
	//response only
	CreatedAt   string `json:"created-at,omitempty"`
	FileURL     string `json:"file-url,omitempty"`
	VersionUrls struct {
		Original string `json:"original"`
	} `json:"version-urls,omitempty"`
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

func NewDocument(accountId string, contactId string, file multipart.File, description string, extension string, label string, mimeType string, public bool) *Document {
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
