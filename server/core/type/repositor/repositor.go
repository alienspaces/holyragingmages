package repositor

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
)

// Repositor -
type Repositor interface {
	Init(p preparer.Preparer, tx *sqlx.Tx) error
	TableName() string
}
