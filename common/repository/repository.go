// Package repository provides methods for interacting with the database
package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Logger -
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

// Repository -
type Repository struct {
	Config       Config
	Log          Logger
	Tx           *sqlx.Tx
	Preparer     *Preparer
	RecordParams map[string]*RecordParam
}

// Config -
type Config struct {
	TableName string
}

// RecordParam -
type RecordParam struct {
	TypeInt        bool
	TypeString     bool
	TypeNullString bool
}

// Init -
func (r *Repository) Init(tx *sqlx.Tx) error {

	r.Log.Info("Initialising repo")

	if r.Tx == nil {
		return errors.New("Tx is nil, cannot initialise")
	}

	return nil
}

// TableName -
func (r *Repository) TableName() string {
	return r.Config.TableName
}

// GetOneRec -
func (r *Repository) GetOneRec(recordID string, rec interface{}) error {

	// preparer
	p := r.Preparer

	// stmt
	stmt := p.GetOneStmt(r)

	r.Log.Info("Get record ID >%s<", recordID)

	err := stmt.QueryRowx(recordID).StructScan(rec)
	if err != nil {
		r.Log.Warn("Failed executing query >%v<", err)
		r.Log.Warn("SQL: >%s<", p.GetOneSQL(r))
		r.Log.Warn("recordID: >%v<", recordID)

		rec = nil

		return err
	}

	r.Log.Info("Record fetched")

	return nil
}

// GetManyRecs -
func (r *Repository) GetManyRecs(
	params map[string]interface{},
	operators map[string]string) (rows *sqlx.Rows, err error) {

	// preparer
	p := r.Preparer

	// stmt
	querySQL := p.GetManySQL(r)

	// tx
	tx := r.Tx

	// params
	querySQL, queryParams, err := r.sqlFromParamsAndOperator(querySQL, params, operators)
	if err != nil {
		r.Log.Info("Failed generating query >%v<", err)
		return nil, err
	}

	r.Log.Info("Query >%s<", querySQL)
	r.Log.Info("Parameters >%+v<", queryParams)

	rows, err = tx.NamedQuery(querySQL, queryParams)
	if err != nil {
		r.Log.Warn("Failed querying row >%v<", err)
		return nil, err
	}

	return rows, nil
}

// CreateRec -
func (r *Repository) CreateRec(rec interface{}) error {

	// preparer
	p := r.Preparer

	// stmt
	stmt := p.CreateStmt(r)

	err := stmt.QueryRowx(rec).StructScan(rec)
	if err != nil {
		r.Log.Warn("Failed executing create >%v<", err)
		return err
	}

	return nil
}

// UpdateOneRec -
func (r *Repository) UpdateOneRec(rec interface{}) error {

	// preparer
	p := r.Preparer

	// stmt
	stmt := p.UpdateOneStmt(r)

	err := stmt.QueryRowx(rec).StructScan(rec)
	if err != nil {
		r.Log.Warn("Failed executing update >%v<", err)
		return err
	}

	return nil
}

// DeleteOne -
func (r *Repository) DeleteOne(id string) error {
	return r.deleteOneRec(id)
}

func (r *Repository) deleteOneRec(recordID string) error {

	params := map[string]interface{}{
		"id":         recordID,
		"deleted_at": NewDeletedAt(),
	}

	// preparer
	p := r.Preparer

	// stmt
	stmt := p.DeleteOneStmt(r)

	res, err := stmt.Exec(params)
	if err != nil {
		r.Log.Warn("Failed executing delete >%v<", err)
		return err
	}

	// rows affected
	raf, err := res.RowsAffected()
	if err != nil {
		r.Log.Warn("Failed executing rows affected >%v<", err)
		return err
	}

	// expect a single row
	if raf != 1 {
		return fmt.Errorf("Expecting to delete exactly one row but deleted >%d<", raf)
	}

	r.Log.Info("Deleted >%d< records", raf)

	return nil
}

// RemoveOne -
func (r *Repository) RemoveOne(id string) error {
	return r.removeOneRec(id)
}

func (r *Repository) removeOneRec(recordID string) error {

	// preparer
	p := r.Preparer

	// stmt
	stmt := p.RemoveOneStmt(r)

	params := map[string]interface{}{
		"id": recordID,
	}

	res, err := stmt.Exec(params)
	if err != nil {
		r.Log.Warn("Failed executing remove >%v<", err)
		return err
	}

	// rows affected
	raf, err := res.RowsAffected()
	if err != nil {
		r.Log.Warn("Failed executing rows affected >%v<", err)
		return err
	}

	// expect a single row
	if raf != 1 {
		return fmt.Errorf("Expecting to remove exactly one row but removed >%d<", raf)
	}

	r.Log.Info("Removed >%d< records", raf)

	return nil
}

// GetOneSQL - Returns SQL
func (r *Repository) GetOneSQL() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE id = $1 AND deleted_at IS NULL", r.TableName())
}

// GetManySQL -
func (r *Repository) GetManySQL() string {
	return fmt.Sprintf("SELECT * FROM %s WHERE deleted_at IS NULL", r.TableName())
}

// CreateSQL -
func (r *Repository) CreateSQL() string {
	return `
INSERT INTO template
	(id, created_at)
VALUES
	(:id, :created_at)
RETURNING *
`
}

// UpdateOneSQL -
func (r *Repository) UpdateOneSQL() string {
	return `
UPDATE template SET
	updated_at = :updated_at
WHERE id 		   = :id
AND   deleted_at IS NULL
RETURNING *
`
}

// UpdateManySQL -
func (r *Repository) UpdateManySQL() string {
	return ``
}

// DeleteOneSQL -
func (r *Repository) DeleteOneSQL() string {
	return fmt.Sprintf("UPDATE %s SET deleted_at = :deleted_at WHERE id = :id AND deleted_at IS NULL RETURNING *", r.TableName())
}

// DeleteManySQL -
func (r *Repository) DeleteManySQL() string {
	return ``
}

// RemoveOneSQL -
func (r *Repository) RemoveOneSQL() string {
	return fmt.Sprintf("DELETE FROM %s WHERE id = :id", r.TableName())
}

// RemoveManySQL -
func (r *Repository) RemoveManySQL() string {
	return ``
}
