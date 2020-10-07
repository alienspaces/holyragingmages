package schema

import "time"

// AccountRequest -
type AccountRequest struct {
	Request
	Data AccountData `json:"data"`
}

// AccountResponse -
type AccountResponse struct {
	Response
	Data []AccountData `json:"data"`
}

// AccountData -
type AccountData struct {
	ID                string    `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Email             string    `json:"email,omitempty"`
	Provider          string    `json:"provider,omitempty"`
	ProviderAccountID string    `json:"provider_account_id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
}

// AccountEntityRequest -
type AccountEntityRequest struct {
	Request
	Data AccountEntityData `json:"data"`
}

// AccountEntityResponse -
type AccountEntityResponse struct {
	Request
	Data []AccountEntityData `json:"data"`
}

// AccountEntityData -
type AccountEntityData struct {
	ID        string    `json:"id,omitempty"`
	AccountID string    `json:"account_id,omitempty"`
	EntityID  string    `json:"entity_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
