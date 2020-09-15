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
	EntityConfig []EntityConfig
}

// EntityConfig -
type EntityConfig struct {
	Record record.Entity
}

// Data -
type Data struct {
	EntityRecs []record.Entity
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

	for _, mageConfig := range t.DataConfig.EntityConfig {

		mageRec, err := t.createEntityRec(mageConfig)
		if err != nil {
			t.Log.Warn("Failed creating mage record >%v<", err)
			return err
		}
		t.Data.EntityRecs = append(t.Data.EntityRecs, mageRec)
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

MAGE_RECS:
	for {
		if len(t.Data.EntityRecs) == 0 {
			break MAGE_RECS
		}
		rec := record.Entity{}
		rec, t.Data.EntityRecs = t.Data.EntityRecs[0], t.Data.EntityRecs[1:]

		err := t.Model.(*model.Model).RemoveEntityRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing mage record >%v<", err)
			return err
		}
	}

	return nil
}
