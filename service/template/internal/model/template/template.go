package template

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/common/model"
)

// Model -
type Model struct {
	model.Model
}

// NewModel -
func NewModel(l model.Logger, tx *sqlx.Tx) (m *Model, err error) {

	m = &Model{
		model.Model{
			Log: l,
			Tx:  tx,
		},
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
