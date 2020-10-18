package harness

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

func (t *Testing) createEntityRec(mageConfig EntityConfig) (record.Entity, error) {

	rec := mageConfig.Record

	// Required properties
	if rec.AccountID == "" {
		accountID, _ := uuid.NewRandom()
		rec.AccountID = accountID.String()
	}
	if rec.Name == "" {
		rec.Name = gofakeit.Name()
	}

	t.Log.Info("Creating testing record >%#v<", rec)

	err := t.Model.(*model.Model).CreateEntityRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing mage record >%v<", err)
		return rec, err
	}
	return rec, nil
}
