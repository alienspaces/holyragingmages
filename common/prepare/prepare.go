package prepare

import (
	"sync"

	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparable"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
)

// Prepare - Methods for preparing and fetching repo statements
type Prepare struct {
	Log logger.Logger
	Tx  *sqlx.Tx
}

var _ preparer.Preparer = &Prepare{}

var getOneSQLList = make(map[string]string)
var getManySQLList = make(map[string]string)
var createSQLList = make(map[string]string)
var updateOneSQLList = make(map[string]string)
var updateManySQLList = make(map[string]string)
var deleteOneSQLList = make(map[string]string)
var deleteManySQLList = make(map[string]string)
var removeOneSQLList = make(map[string]string)
var removeManySQLList = make(map[string]string)

var getOneStmtList = make(map[string]*sqlx.Stmt)
var getManyStmtList = make(map[string]*sqlx.NamedStmt)
var createStmtList = make(map[string]*sqlx.NamedStmt)
var updateOneStmtList = make(map[string]*sqlx.NamedStmt)
var updateManyStmtList = make(map[string]*sqlx.NamedStmt)
var deleteOneStmtList = make(map[string]*sqlx.NamedStmt)
var deleteManyStmtList = make(map[string]*sqlx.NamedStmt)
var removeOneStmtList = make(map[string]*sqlx.NamedStmt)
var removeManyStmtList = make(map[string]*sqlx.NamedStmt)

// prepared
var prepared = make(map[string]bool)

// mutex
var mutex = &sync.Mutex{}

// NewPrepare -
func NewPrepare(l logger.Logger, tx *sqlx.Tx) (*Prepare, error) {

	p := Prepare{
		Log: l,
		Tx:  tx,
	}

	return &p, nil
}

// Prepare - Prepares all repo SQL statements for faster execution
func (p *Prepare) Prepare(m preparable.Preparable) error {

	// lock/unlock
	mutex.Lock()
	defer mutex.Unlock()

	// already prepared
	if _, ok := prepared[m.TableName()]; ok {
		return nil
	}

	p.Log.Info("Preparing statements >%s<", m.TableName())

	// get by id
	query := m.GetOneSQL()

	getOneStmt, err := p.Tx.Preparex(query)
	if err != nil {
		p.Log.Warn("Error preparing GetOneSQL statement >%v<", err)
		return err
	}

	getOneSQLList[m.TableName()] = query
	getOneStmtList[m.TableName()] = getOneStmt

	// get many
	query = m.GetManySQL()

	getManyStmt, err := p.Tx.PrepareNamed(m.GetManySQL())
	if err != nil {
		p.Log.Warn("Error preparing GetManySQL statement >%v<", err)
		return err
	}

	getManySQLList[m.TableName()] = query
	getManyStmtList[m.TableName()] = getManyStmt

	// create
	query = m.CreateSQL()

	createStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing CreateSQL statement >%v<", err)
		return err
	}

	createSQLList[m.TableName()] = query
	createStmtList[m.TableName()] = createStmt

	// update
	query = m.UpdateOneSQL()

	updateOneStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing UpdateOneSQL statement >%v<", err)
		return err
	}

	updateOneSQLList[m.TableName()] = query
	updateOneStmtList[m.TableName()] = updateOneStmt

	// update many
	query = m.UpdateManySQL()

	updateManyStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing UpdateManySQL statement >%v<", err)
		return err
	}

	updateManySQLList[m.TableName()] = query
	updateManyStmtList[m.TableName()] = updateManyStmt

	// delete
	query = m.DeleteOneSQL()

	deleteStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing DeleteSQL statement >%v<", err)
		return err
	}

	deleteOneSQLList[m.TableName()] = query
	deleteOneStmtList[m.TableName()] = deleteStmt

	// delete many
	query = m.DeleteManySQL()

	deleteManyStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing DeleteManySQL statement >%v<", err)
		return err
	}

	deleteManySQLList[m.TableName()] = query
	deleteManyStmtList[m.TableName()] = deleteManyStmt

	// remove
	query = m.RemoveOneSQL()

	removeStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing RemoveSQL statement >%v<", err)
		return err
	}

	removeOneSQLList[m.TableName()] = query
	removeOneStmtList[m.TableName()] = removeStmt

	// remove many
	query = m.RemoveManySQL()

	removeManyStmt, err := p.Tx.PrepareNamed(query)
	if err != nil {
		p.Log.Warn("Error preparing RemoveManySQL statement >%v<", err)
		return err
	}

	removeManySQLList[m.TableName()] = query
	removeManyStmtList[m.TableName()] = removeManyStmt

	prepared[m.TableName()] = true

	return nil
}

// GetOneStmt -
func (p *Prepare) GetOneStmt(m preparable.Preparable) *sqlx.Stmt {

	stmt := getOneStmtList[m.TableName()]

	return p.Tx.Stmtx(stmt)
}

// GetManyStmt -
func (p *Prepare) GetManyStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := getManyStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// CreateStmt -
func (p *Prepare) CreateStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := createStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// UpdateOneStmt -
func (p *Prepare) UpdateOneStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := updateOneStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// UpdateManyStmt -
func (p *Prepare) UpdateManyStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := updateManyStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// DeleteOneStmt -
func (p *Prepare) DeleteOneStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := deleteOneStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// DeleteManyStmt -
func (p *Prepare) DeleteManyStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := deleteManyStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// RemoveOneStmt -
func (p *Prepare) RemoveOneStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := removeOneStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// RemoveManyStmt -
func (p *Prepare) RemoveManyStmt(m preparable.Preparable) *sqlx.NamedStmt {

	stmt := removeManyStmtList[m.TableName()]

	return p.Tx.NamedStmt(stmt)
}

// GetOneSQL -
func (p *Prepare) GetOneSQL(m preparable.Preparable) string {

	query := getOneSQLList[m.TableName()]

	return query
}

// GetManySQL -
func (p *Prepare) GetManySQL(m preparable.Preparable) string {

	query := getManySQLList[m.TableName()]

	return query
}

// CreateSQL -
func (p *Prepare) CreateSQL(m preparable.Preparable) string {

	query := createSQLList[m.TableName()]

	return query
}

// UpdateOneSQL -
func (p *Prepare) UpdateOneSQL(m preparable.Preparable) string {

	query := updateOneSQLList[m.TableName()]

	return query
}

// UpdateManySQL -
func (p *Prepare) UpdateManySQL(m preparable.Preparable) string {

	query := updateManySQLList[m.TableName()]

	return query
}

// DeleteOneSQL -
func (p *Prepare) DeleteOneSQL(m preparable.Preparable) string {

	query := deleteOneSQLList[m.TableName()]

	return query
}

// DeleteManySQL -
func (p *Prepare) DeleteManySQL(m preparable.Preparable) string {

	query := deleteManySQLList[m.TableName()]

	return query
}

// RemoveOneSQL -
func (p *Prepare) RemoveOneSQL(m preparable.Preparable) string {

	query := removeOneSQLList[m.TableName()]

	return query
}

// RemoveManySQL -
func (p *Prepare) RemoveManySQL(m preparable.Preparable) string {

	query := removeManySQLList[m.TableName()]

	return query
}
