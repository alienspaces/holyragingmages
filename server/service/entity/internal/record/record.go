package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// EntityType -
const (
	EntityTypeMage             string = "mage"
	EntityTypeFamilliar        string = "familliar"
	EntityTypePlayerMage       string = "player-mage"
	EntityTypePlayerFamilliar  string = "player-familliar"
	EntityTypeStarterMage      string = "starter-mage"
	EntityTypeStarterFamilliar string = "starter-familliar"
)

// MageAvatar -
const (
	MageAvatarDarkArmoured         string = "dark-armoured"
	MageAvatarRedStripeDruid       string = "red-stripe-druid"
	MageAvatarRedFairy             string = "red-fairy"
	MageAvatarRedStripeNecromancer string = "red-stripe-necromancer"
	MageAvatarGreenElven           string = "green-elven"
)

// FamilliarAvatar -
const (
	FamilliarAvatarBrownCyclopsBat      string = "brown-cyclops-bat"
	FamilliarAvatarBrownYeti            string = "brown-yeti"
	FamilliarAvatarGreenTribble         string = "green-tribble"
	FamilliarAvatarGreyCyclops          string = "grey-cyclops"
	FamilliarAvatarOrangeSpottedTribble string = "orange-spotted-tribble"
	FamilliarAvatarPurpleBat            string = "purple-bat"
	FamilliarAvatarPurpleMinotaur       string = "purple-minotaur"
)

// Entity -
type Entity struct {
	repository.Record
	EntityType       string `db:"entity_type"`
	Name             string `db:"name"`
	Avatar           string `db:"avatar"`
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
