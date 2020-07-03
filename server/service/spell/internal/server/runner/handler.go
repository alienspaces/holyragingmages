package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/record"
)

// GetSpellsHandler -
func (rnr *Runner) GetSpellsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get spells handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Spell
	var err error

	id := pp.ByName("spell_id")

	// single resource
	if id != "" {

		l.Info("Getting spell record ID >%s<", id)

		rec, err := m.(*model.Model).GetSpellRec(id, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		// resource not found
		if rec == nil {
			rnr.WriteNotFoundError(l, w, id)
			return
		}

		recs = append(recs, rec)

	} else {

		l.Info("Querying spell records")

		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetSpellRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.SpellData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToSpellResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.SpellResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostSpellsHandler -
func (rnr *Runner) PostSpellsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post spells handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("spell_id")

	req := schema.SpellRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Spell{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.SpellRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateSpellRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToSpellResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.SpellResponse{
		Data: []schema.SpellData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutSpellsHandler -
func (rnr *Runner) PutSpellsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put spells handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("spell_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetSpellRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.SpellRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.SpellRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateSpellRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToSpellResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.SpellResponse{
		Data: []schema.SpellData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// SpellRequestDataToRecord -
func (rnr *Runner) SpellRequestDataToRecord(data schema.SpellData, rec *record.Spell) error {

	rec.Name = data.Name
	rec.Description = data.Description

	return nil
}

// RecordToSpellResponseData -
func (rnr *Runner) RecordToSpellResponseData(spellRec *record.Spell) (schema.SpellData, error) {

	data := schema.SpellData{
		ID:          spellRec.ID,
		Name:        spellRec.Name,
		Description: spellRec.Description,
		CreatedAt:   spellRec.CreatedAt,
		UpdatedAt:   spellRec.UpdatedAt.Time,
	}

	return data, nil
}
