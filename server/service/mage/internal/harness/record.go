package harness

import (
	"github.com/brianvoe/gofakeit/v5"

	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/record"
)

func (t *Testing) createMageRec(mageConfig MageConfig) (record.Mage, error) {

	rec := mageConfig.Record

	// Required properties
	if rec.Name == "" {
		rec.Name = gofakeit.Name()
	}

	t.Log.Info("Creating testing record >%#v<", rec)

	err := t.Model.(*model.Model).CreateMageRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing mage record >%v<", err)
		return rec, err
	}
	return rec, nil
}
