package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/record"
)

func (t *Testing) createItemRec(itemConfig ItemConfig) (record.Item, error) {

	rec := itemConfig.Record

	t.Log.Info("Creating testing record >%#v<", rec)

	err := t.Model.(*model.Model).CreateItemRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing item record >%v<", err)
		return rec, err
	}

	return rec, nil
}
