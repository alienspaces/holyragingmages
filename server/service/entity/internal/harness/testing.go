package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/harness"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	EntityConfig        []EntityConfig
	AccountEntityConfig []AccountEntityConfig
}

// AccountEntityConfig -
type AccountEntityConfig struct {
	Record       record.AccountEntity
	EntityConfig []EntityConfig
}

// EntityConfig -
type EntityConfig struct {
	Record record.Entity
}

// Data -
type Data struct {
	EntityRecs        []record.Entity
	AccountEntityRecs []record.AccountEntity
}

// NewTesting -
func NewTesting(config DataConfig) (t *Testing, err error) {

	// harness
	t = &Testing{}

	// modeller
	t.ModellerFunc = t.Modeller

	// data
	t.CreateDataFunc = t.CreateData
	t.RemoveDataFunc = t.RemoveData

	t.DataConfig = config
	t.Data = &Data{}

	err = t.Init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Modeller -
func (t *Testing) Modeller() (modeller.Modeller, error) {

	m, err := model.NewModel(t.Config, t.Log, t.Store)
	if err != nil {
		t.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// CreateData - Custom data
func (t *Testing) CreateData() error {

	// Account entity with entity records
	for _, accountEntityConfig := range t.DataConfig.AccountEntityConfig {

		// NOTE: For an account entity record to be created there must
		// be at least one entity record created.
		if len(accountEntityConfig.EntityConfig) == 0 {
			accountEntityConfig.EntityConfig = []EntityConfig{
				{
					Record: record.Entity{},
				},
			}
		}

		for _, entityConfig := range accountEntityConfig.EntityConfig {

			entityRec, err := t.createEntityRec(entityConfig)
			if err != nil {
				t.Log.Warn("Failed creating entity record >%v<", err)
				return err
			}
			t.Data.EntityRecs = append(t.Data.EntityRecs, entityRec)

			accountEntityRec, err := t.createAccountEntityRec(entityRec, accountEntityConfig)
			if err != nil {
				t.Log.Warn("Failed creating account entity record >%v<", err)
				return err
			}
			t.Data.AccountEntityRecs = append(t.Data.AccountEntityRecs, accountEntityRec)
		}
	}

	// Stand alone entity records
	for _, entityConfig := range t.DataConfig.EntityConfig {

		entityRec, err := t.createEntityRec(entityConfig)
		if err != nil {
			t.Log.Warn("Failed creating entity record >%v<", err)
			return err
		}
		t.Data.EntityRecs = append(t.Data.EntityRecs, entityRec)
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

ACCOUNT_RECS:
	for {
		if len(t.Data.AccountEntityRecs) == 0 {
			break ACCOUNT_RECS
		}
		rec := record.AccountEntity{}
		rec, t.Data.AccountEntityRecs = t.Data.AccountEntityRecs[0], t.Data.AccountEntityRecs[1:]

		err := t.Model.(*model.Model).RemoveAccountEntityRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing account entity record >%v<", err)
			return err
		}
	}

ENTITY_RECS:
	for {
		if len(t.Data.EntityRecs) == 0 {
			break ENTITY_RECS
		}
		rec := record.Entity{}
		rec, t.Data.EntityRecs = t.Data.EntityRecs[0], t.Data.EntityRecs[1:]

		err := t.Model.(*model.Model).RemoveEntityRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing entity record >%v<", err)
			return err
		}
	}

	return nil
}
