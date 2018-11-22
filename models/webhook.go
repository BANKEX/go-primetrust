package models

import "time"

const (
	WebhookType = "contacts"
)

type WebhookResponse struct {
	CollectionResponse
	Data WebhookData `json:"data"`
}

type WebhookAttribute struct {
	AccountID string `json:"account_id"`
	URL       string `json:"url"`
	Enabled   bool   `json:"enabled"`
	//only for response
	ContactEmail string    `json:"contact-email"`
	CreatedAt    time.Time `json:"created-at"`
	Failures     int       `json:"failures"`
	SharedSecret string    `json:"shared-secret"`
	UpdatedAt    time.Time `json:"updated-at"`
}

type WebhookData struct {
	Attributes WebhookAttribute `json:"attributes"`
	Type       string           `json:"type"`
	//only for response
	Links         Links         `json:"links"`
	Relationships Relationships `json:"relationships"`
	ID            string        `json:"id,omitempty"`
}

type Webhook struct {
	Data WebhookData `json:"data"`
}

func NewWebhook(accountId string, url string, enabled bool) *Webhook {
	webhook := Webhook{
		Data: WebhookData{
			Type: WebhookType,
			Attributes: WebhookAttribute{
				Enabled:   enabled,
				URL:       url,
				AccountID: accountId,
			},
		},
	}
	return &webhook
}
