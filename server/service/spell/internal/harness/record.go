package harness

import (
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/record"
)

func (t *Testing) createSpellRec(spellConfig SpellConfig) (record.Spell, error) {

	rec := spellConfig.Record

	t.Log.Info("Creating testing record >%#v<", rec)

	err := t.Model.(*model.Model).CreateSpellRec(&rec)
	if err != nil {
		t.Log.Warn("Failed creating testing spell record >%v<", err)
		return rec, err
	}
	return rec, nil
}
