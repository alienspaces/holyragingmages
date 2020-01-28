package template

// TODO: Complete baseline template example model
// - Add proper initialisation function with new DB tx
// - Reinitialise repostore when given a new DB tx
// - Add Query example

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/model"
)

// Model -
type Model struct {
	Log model.Logger
	Tx  *sqlx.Tx
}

// NewModel - Returns a new model
func NewModel(l model.Logger, tx *sqlx.Tx) (m *Model, err error) {

	m = &Model{
		Log: l,
		Tx:  tx,
	}

	// initialize only if we have a tx
	if tx != nil {
		err = m.Init(tx)
		if err != nil {
			m.Log.Printf("Failed model init >%v<", err)
			return nil, err
		}
	}

	return m, nil
}

// Init -
func (m *Model) Init(tx *sqlx.Tx) (err error) {

	// tx required
	if tx == nil {
		m.Log.Printf("Failed init, tx is required")
		return fmt.Errorf("Failed init, tx is required")
	}

	return nil
}
