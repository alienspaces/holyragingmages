package schema

import (
	"time"
)

// EntityResponse -
type EntityResponse struct {
	Response
	Data []EntityData `json:"data"`
}

// EntityRequest -
type EntityRequest struct {
	Request
	Data EntityData `json:"data"`
}

// EntityData -
type EntityData struct {
	ID               string    `json:"id,omitempty"`
	Name             string    `json:"name"`
	Strength         int       `json:"strength"`
	Dexterity        int       `json:"dexterity"`
	Intelligence     int       `json:"intelligence"`
	AttributePoints  int64     `json:"attribute_points"`
	ExperiencePoints int64     `json:"experience_points"`
	Coins            int64     `json:"coins"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}
