package repositor

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
)

// Repositor -
type Repositor interface {
	Init(p preparer.Preparer, tx *sqlx.Tx) error
	TableName() string
}
