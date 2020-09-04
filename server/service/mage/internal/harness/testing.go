package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/harness"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	MageConfig []MageConfig
}

// MageConfig -
type MageConfig struct {
	Record record.Mage
}

// Data -
type Data struct {
	MageRecs []record.Mage
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

	for _, mageConfig := range t.DataConfig.MageConfig {

		mageRec, err := t.createMageRec(mageConfig)
		if err != nil {
			t.Log.Warn("Failed creating mage record >%v<", err)
			return err
		}
		t.Data.MageRecs = append(t.Data.MageRecs, mageRec)
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

MAGE_RECS:
	for {
		if len(t.Data.MageRecs) == 0 {
			break MAGE_RECS
		}
		rec := record.Mage{}
		rec, t.Data.MageRecs = t.Data.MageRecs[0], t.Data.MageRecs[1:]

		err := t.Model.(*model.Model).RemoveMageRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing mage record >%v<", err)
			return err
		}
	}

	return nil
}
