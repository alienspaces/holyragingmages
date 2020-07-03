package schema

import "time"

// ItemResponse -
type ItemResponse struct {
	Response
	Data []ItemData `json:"data"`
}

// ItemRequest -
type ItemRequest struct {
	Request
	Data ItemData `json:"data"`
}

// ItemData -
type ItemData struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
