package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/harness"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	ItemConfig []ItemConfig
}

// ItemConfig -
type ItemConfig struct {
	Record record.Item
}

// Data -
type Data struct {
	ItemRecs []record.Item
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

	for _, itemConfig := range t.DataConfig.ItemConfig {

		itemRec, err := t.createItemRec(itemConfig)
		if err != nil {
			t.Log.Warn("Failed creating item record >%v<", err)
			return err
		}
		t.Data.ItemRecs = append(t.Data.ItemRecs, itemRec)
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

ITEM_RECS:
	for {
		if len(t.Data.ItemRecs) == 0 {
			break ITEM_RECS
		}
		rec := record.Item{}
		rec, t.Data.ItemRecs = t.Data.ItemRecs[0], t.Data.ItemRecs[1:]

		err := t.Model.(*model.Model).RemoveItemRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing item record >%v<", err)
			return err
		}
	}

	return nil
}
