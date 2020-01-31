package model

import (
	"gitlab.com/alienspaces/holyragingmages/common/model"
)

// Model -
type Model struct {
	model.Model
}

// NewModel -
func NewModel(c model.Configurer, l model.Logger, s model.Storer) (m *Model, err error) {

	m = &Model{
		model.Model{
			Config: c,
			Log:    l,
			Store:  s,
		},
	}

	return m, nil
}
