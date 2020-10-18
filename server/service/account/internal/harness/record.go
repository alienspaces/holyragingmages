package harness

import (
	"github.com/brianvoe/gofakeit"

	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

func (t *Testing) createAccountRec(accountConfig AccountConfig) (record.Account, error) {

	rec := accountConfig.Record

	t.Log.Info("Creating test account record >%#v<", rec)

	// NOTE: Add default values for required properties here
	if rec.Name == "" {
		rec.Name = gofakeit.Name()
	}

	if rec.Email == "" {
		rec.Email = gofakeit.Email()
	}

	if rec.Provider == "" {
		rec.Provider = record.AccountProviderGoogle
	}

	if rec.ProviderAccountID == "" {
		rec.ProviderAccountID = gofakeit.UUID()
	}

	err := t.Model.(*model.Model).CreateAccountRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing account record >%v<", err)
		return rec, err
	}
	return rec, nil
}

func (t *Testing) createAccountRoleRec(accountRec record.Account, accountRoleConfig AccountRoleConfig) (record.AccountRole, error) {

	rec := accountRoleConfig.Record

	t.Log.Info("Creating test account role record >%#v<", rec)

	// NOTE: Add default values for required properties here
	rec.AccountID = accountRec.ID

	if rec.Role == "" {
		rec.Role = record.AccountRoleDefault
	}

	err := t.Model.(*model.Model).CreateAccountRoleRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing account role record >%v<", err)
		return rec, err
	}
	return rec, nil
}
