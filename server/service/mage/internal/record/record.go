package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// Mage -
type Mage struct {
	repository.Record
	Name         string `db:"name"`
	Strength     int    `db:"strength"`
	Dexterity    int    `db:"dexterity"`
	Intelligence int    `db:"intelligence"`
	Experience   int64  `db:"experience"`
	Coin         int64  `db:"coin"`
}
