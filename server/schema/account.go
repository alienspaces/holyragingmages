package schema

import "time"

// AccountResponse -
type AccountResponse struct {
	Response
	Data []AccountData `json:"data"`
}

// AccountRequest -
type AccountRequest struct {
	Request
	Data AccountData `json:"data"`
}

// AccountData -
type AccountData struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
