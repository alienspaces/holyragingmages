package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/core/harness"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	AccountConfig []AccountConfig
}

// AccountConfig -
type AccountConfig struct {
	Record            record.Account
	AccountRoleConfig []AccountRoleConfig
}

// AccountRoleConfig -
type AccountRoleConfig struct {
	Record record.AccountRole
}

// Data -
type Data struct {
	AccountRecs     []record.Account
	AccountRoleRecs []record.AccountRole
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

	for _, accountConfig := range t.DataConfig.AccountConfig {

		accountRec, err := t.createAccountRec(accountConfig)
		if err != nil {
			t.Log.Warn("Failed creating account record >%v<", err)
			return err
		}
		t.Data.AccountRecs = append(t.Data.AccountRecs, accountRec)

		for _, accountRoleConfig := range accountConfig.AccountRoleConfig {
			accountRoleRec, err := t.createAccountRoleRec(accountRec, accountRoleConfig)
			if err != nil {
				t.Log.Warn("Failed creating account role record >%v<", err)
				return err
			}
			t.Data.AccountRoleRecs = append(t.Data.AccountRoleRecs, accountRoleRec)
		}
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

ACCOUNT_ROLE_RECS:
	for {
		if len(t.Data.AccountRoleRecs) == 0 {
			break ACCOUNT_ROLE_RECS
		}
		rec := record.AccountRole{}
		rec, t.Data.AccountRoleRecs = t.Data.AccountRoleRecs[0], t.Data.AccountRoleRecs[1:]

		err := t.Model.(*model.Model).RemoveAccountRoleRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing account role record >%v<", err)
			return err
		}
	}

ACCOUNT_RECS:
	for {
		if len(t.Data.AccountRecs) == 0 {
			break ACCOUNT_RECS
		}
		rec := record.Account{}
		rec, t.Data.AccountRecs = t.Data.AccountRecs[0], t.Data.AccountRecs[1:]

		err := t.Model.(*model.Model).RemoveAccountRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing account record >%v<", err)
			return err
		}
	}

	return nil
}
