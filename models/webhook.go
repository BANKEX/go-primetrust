package models

import "time"

const (
	WebhooksType = "webhook-configs"
)

type WebhooksResponse struct {
	CollectionResponse
	Data []WebhookData `json:"data"`
}

type WebhookAttribute struct {
	AccountID    string    `json:"account_id"`
	URL          string    `json:"url"`
	Enabled      bool      `json:"enabled"`
	ContactEmail string    `json:"contact-email"`
	SharedSecret string    `json:"shared-secret"`
	CreatedAt    time.Time `json:"created-at,omitempty"`
	Failures     int       `json:"failures,omitempty"`
	UpdatedAt    time.Time `json:"updated-at,omitempty"`
}

type WebhookData struct {
	ID            string           `json:"id,omitempty"`
	Type          string           `json:"type"`
	Attributes    WebhookAttribute `json:"attributes"`
	Links         Links            `json:"links,omitempty"`
	Relationships Relationships    `json:"relationships,omitempty"`
}

type Webhook struct {
	Data WebhookData `json:"data"`
}

func NewWebhook(accountId string, url string, email string, secret string, enabled bool) *Webhook {
	webhook := Webhook{
		Data: WebhookData{
			Type: WebhooksType,
			Attributes: WebhookAttribute{
				Enabled:      enabled,
				URL:          url,
				AccountID:    accountId,
				ContactEmail: email,
				SharedSecret: secret,
			},
		},
	}
	return &webhook
}

type WebhookPayloadData struct {
	Changes    []string          `json:"changes,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type WebhookPayload struct {
	ID               string             `json:"id,omitempty"`
	Action           string             `json:"action,omitempty"`
	CreatedAt        time.Time          `json:"created-at,omitempty"`
	Data             WebhookPayloadData `json:"data,omitempty"`
	Failures         int                `json:"failures,omitempty"`
	LastRequestAt    time.Time          `json:"last-request-at,omitempty"`
	LastResponseCode int                `json:"last-response-code,omitempty"`
	ResourceID       string             `json:"resource-id,omitempty"`
	ResourceType     string             `json:"resource-type,omitempty"`
	Success          bool               `json:"success,omitempty"`
}
