package record

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/repository"
)

// Item -
type Item struct {
	repository.Record
	Name        string `db:"name"`
	Description string `db:"description"`
}
