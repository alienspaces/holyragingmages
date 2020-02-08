package template

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

type TestData struct {
	TemplateRec *record.Template
}

func (td *TestData) Setup(tx *sqlx.Tx) error {

	return nil
}

func (td *TestData) Teardown(tx *sqlx.Tx) error {

	return nil
}
