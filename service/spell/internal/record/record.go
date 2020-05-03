package record

import (
	"gitlab.com/alienspaces/holyragingmages/common/repository"
)

// Spell -
type Spell struct {
	repository.Record
	Name        string `db:"name"`
	Description string `db:"description"`
}
