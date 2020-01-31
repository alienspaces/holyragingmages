package model

import (
	"gitlab.com/alienspaces/holyragingmages/common/model"
	"gitlab.com/alienspaces/holyragingmages/common/service"
)

// Model -
type Model struct {
	model.Model
}

// NewModel -
func NewModel(c service.Configurer, l service.Logger, s service.Storer) (*Model, error) {

	m := &Model{
		model.Model{
			Config: c,
			Log:    l,
			Store:  s,
		},
	}

	return m, nil
}
