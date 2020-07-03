package schema

import "time"

// SpellResponse -
type SpellResponse struct {
	Response
	Data []SpellData `json:"data"`
}

// SpellRequest -
type SpellRequest struct {
	Request
	Data SpellData `json:"data"`
}

// SpellData -
type SpellData struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
