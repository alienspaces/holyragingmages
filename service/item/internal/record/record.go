package record

import (
	"gitlab.com/alienspaces/holyragingmages/common/repository"
)

// Item -
type Item struct {
	repository.Record
	Name        string `db:"name"`
	Description string `db:"description"`
}
