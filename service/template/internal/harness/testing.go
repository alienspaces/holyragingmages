package harness

import (
	"gitlab.com/alienspaces/holyragingmages/common/harness"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig *DataConfig
}

// DataConfig -
type DataConfig struct {
	harness.DataConfig
	TemplateConfig []TemplateConfig
}

// TemplateConfig -
type TemplateConfig struct {
	Count  int
	Record *record.Template
}

// Data -
type Data struct {
	harness.Data
	TemplateRecs []*record.Template
}

// NewTesting -
func NewTesting() (t *Testing, err error) {

	// harness
	t = &Testing{}

	// modeller
	t.ModellerFunc = t.Modeller

	// data
	t.CreateDataFunc = t.CreateData
	t.RemoveDataFunc = t.RemoveData

	t.Data = &Data{}

	return t, nil
}

// Modeller -
func (t *Testing) Modeller() (modeller.Modeller, error) {

	t.Log.Info("** Testing Model **")

	m, err := model.NewModel(t.Config, t.Log, t.Store)
	if err != nil {
		t.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// CreateData - Custom data
func (t *Testing) CreateData() error {

	r := t.Model.(*model.Model).TemplateRepository()

	rec := r.NewRecord()

	t.Log.Warn("Test record >%#v<", rec)

	err := r.CreateTestRecord(rec)
	if err != nil {
		t.Log.Warn("Failed creating test template record >%v<", err)
		return err
	}

	t.Data.TemplateRecs = append(t.Data.TemplateRecs, rec)

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

	return nil
}
