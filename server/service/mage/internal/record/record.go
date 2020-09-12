package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// Mage -
type Mage struct {
	repository.Record
	Name             string `db:"name"`
	Strength         int    `db:"strength"`
	Dexterity        int    `db:"dexterity"`
	Intelligence     int    `db:"intelligence"`
	AttributePoints  int64  `db:"attribute_points"`
	ExperiencePoints int64  `db:"experience_points"`
	Coins            int64  `db:"coins"`
}
