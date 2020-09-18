package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

func (t *Testing) createAccountRec(accountConfig AccountConfig) (record.Account, error) {

	rec := accountConfig.Record

	t.Log.Info("Creating testing record >%#v<", rec)

	err := t.Model.(*model.Model).CreateAccountRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing account record >%v<", err)
		return rec, err
	}
	return rec, nil
}
