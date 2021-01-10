package harness

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"

	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

func (t *Testing) createEntityRec(entityConfig EntityConfig) (record.Entity, error) {

	entityRec := entityConfig.Record

	if entityRec.EntityType == "" {
		entityRec.EntityType = record.EntityTypePlayerCharacter
	}

	if entityRec.Name == "" {
		entityRec.Name = gofakeit.Name()
	}

	t.Log.Info("Creating entity testing record >%#v<", entityRec)

	err := t.Model.(*model.Model).CreateEntityRec(&entityRec)
	if err != nil {
		t.Log.Warn("Failed creating testing entity record >%v<", err)
		return entityRec, err
	}

	return entityRec, nil
}

func (t *Testing) createAccountEntityRec(entityRec record.Entity, accountEntityConfig AccountEntityConfig) (record.AccountEntity, error) {

	accountEntityRec := accountEntityConfig.Record

	t.Log.Info("Creating account entity testing record >%#v<", accountEntityRec)

	if accountEntityRec.AccountID == "" {
		accountID, _ := uuid.NewRandom()
		accountEntityRec.AccountID = accountID.String()
	}
	accountEntityRec.EntityID = entityRec.ID

	err := t.Model.(*model.Model).CreateAccountEntityRec(&accountEntityRec)
	if err != nil {
		t.Log.Warn("Failed creating testing account entity record >%v<", err)
		return accountEntityRec, err
	}

	return accountEntityRec, nil
}
