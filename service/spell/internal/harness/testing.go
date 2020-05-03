package harness

import (
	"gitlab.com/alienspaces/holyragingmages/common/harness"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	SpellConfig []SpellConfig
}

// SpellConfig -
type SpellConfig struct {
	Count  int
	Record *record.Spell
}

// Data -
type Data struct {
	SpellRecs []record.Spell
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

	rec := record.Spell{}

	err := t.Model.(*model.Model).CreateSpellRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing spell record >%v<", err)
		return err
	}

	t.Data.SpellRecs = append(t.Data.SpellRecs, rec)

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

SPELL_RECS:
	for {
		if len(t.Data.SpellRecs) == 0 {
			break SPELL_RECS
		}
		rec := record.Spell{}
		rec, t.Data.SpellRecs = t.Data.SpellRecs[0], t.Data.SpellRecs[1:]

		err := t.Model.(*model.Model).RemoveSpellRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing spell record >%v<", err)
			return err
		}
	}

	return nil
}
