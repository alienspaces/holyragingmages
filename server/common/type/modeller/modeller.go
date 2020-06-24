package modeller

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/common/type/preparer"
)

// Modeller -
type Modeller interface {
	Init(p preparer.Preparer, tx *sqlx.Tx) (err error)
}
