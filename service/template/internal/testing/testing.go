package testing

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// Data -
type Data struct {
	TemplateRecs *record.Template
}

// Setup -
func (d *Data) Setup(tx *sqlx.Tx) error {

	return nil
}

// Teardown -
func (d *Data) Teardown(tx *sqlx.Tx) error {

	return nil
}
