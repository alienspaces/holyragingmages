package preparer

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/preparable"
)

// Preparer -
type Preparer interface {
	Prepare(m preparable.Preparable) error
	GetOneStmt(m preparable.Preparable) *sqlx.Stmt
	GetManyStmt(m preparable.Preparable) *sqlx.NamedStmt
	CreateStmt(m preparable.Preparable) *sqlx.NamedStmt
	UpdateOneStmt(m preparable.Preparable) *sqlx.NamedStmt
	UpdateManyStmt(m preparable.Preparable) *sqlx.NamedStmt
	DeleteOneStmt(m preparable.Preparable) *sqlx.NamedStmt
	DeleteManyStmt(m preparable.Preparable) *sqlx.NamedStmt
	RemoveOneStmt(m preparable.Preparable) *sqlx.NamedStmt
	RemoveManyStmt(m preparable.Preparable) *sqlx.NamedStmt
	GetOneSQL(m preparable.Preparable) string
	GetManySQL(m preparable.Preparable) string
	CreateSQL(m preparable.Preparable) string
	UpdateOneSQL(m preparable.Preparable) string
	UpdateManySQL(m preparable.Preparable) string
	DeleteOneSQL(m preparable.Preparable) string
	DeleteManySQL(m preparable.Preparable) string
	RemoveOneSQL(m preparable.Preparable) string
	RemoveManySQL(m preparable.Preparable) string
}
