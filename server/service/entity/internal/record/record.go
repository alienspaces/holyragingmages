package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// EntityType -
const (
	EntityTypeFamilliar          string = "familliar"
	EntityTypeNonPlayerCharacter string = "non-player"
	EntityTypePlayerCharacter    string = "player"
	EntityTypeStarterCharacter   string = "starter"
)

// Entity -
type Entity struct {
	repository.Record
	EntityType       string `db:"entity_type"`
	Name             string `db:"name"`
	Strength         int    `db:"strength"`
	Dexterity        int    `db:"dexterity"`
	Intelligence     int    `db:"intelligence"`
	AttributePoints  int64  `db:"attribute_points"`
	ExperiencePoints int64  `db:"experience_points"`
	Coins            int64  `db:"coins"`
}

// AccountEntity -
type AccountEntity struct {
	repository.Record
	AccountID string `db:"account_id"`
	EntityID  string `db:"entity_id"`
}

// EntityItem -
type EntityItem struct {
	repository.Record
	ItemID   string `db:"item_id"`
	EntityID string `db:"entity_id"`
}

// EntitySpell -
type EntitySpell struct {
	repository.Record
	SpellID  string `db:"spell_id"`
	EntityID string `db:"entity_id"`
}
