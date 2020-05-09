package runner

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/record"
)

// SpellResponse -
type SpellResponse struct {
	server.Response
	Data []SpellData `json:"data"`
}

// SpellRequest -
type SpellRequest struct {
	server.Request
	Data SpellData `json:"data"`
}

// SpellData -
type SpellData struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// GetSpellsHandler -
func (rnr *Runner) GetSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get spells handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Spell
	var err error

	id := p.ByName("spell_id")

	// single resource
	if id != "" {

		rnr.Log.Info("Getting spell ID >%s<", id)

		rec, err := m.(*model.Model).GetSpellRec(id, false)
		if err != nil {
			rnr.WriteModelError(w, err)
			return
		}

		// resource not found
		if rec == nil {
			rnr.WriteNotFoundError(w, id)
			return
		}

		recs = append(recs, rec)

	} else {

		rnr.Log.Info("Getting all spell records")

		recs, err = m.(*model.Model).GetSpellRecs(nil, nil, false)
		if err != nil {
			rnr.WriteModelError(w, err)
			return
		}
	}

	// assign response properties
	data := []SpellData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToSpellResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(w, err)
			return
		}

		data = append(data, responseData)
	}

	res := SpellResponse{
		Data: data,
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostSpellsHandler -
func (rnr *Runner) PostSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post spells handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("spell_id")

	req := SpellRequest{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	rec := record.Spell{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.SpellRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).CreateSpellRec(&rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToSpellResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := SpellResponse{
		Data: []SpellData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutSpellsHandler -
func (rnr *Runner) PutSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Put spells handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("spell_id")

	rnr.Log.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetSpellRec(id, false)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(w, id)
		return
	}

	req := SpellRequest{}

	err = rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// record data
	err = rnr.SpellRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).UpdateSpellRec(rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToSpellResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := SpellResponse{
		Data: []SpellData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// SpellRequestDataToRecord -
func (rnr *Runner) SpellRequestDataToRecord(data SpellData, rec *record.Spell) error {

	rec.Name = data.Name
	rec.Description = data.Description

	return nil
}

// RecordToSpellResponseData -
func (rnr *Runner) RecordToSpellResponseData(spellRec *record.Spell) (SpellData, error) {

	data := SpellData{
		ID:          spellRec.ID,
		Name:        spellRec.Name,
		Description: spellRec.Description,
		CreatedAt:   spellRec.CreatedAt,
		UpdatedAt:   spellRec.UpdatedAt.Time,
	}

	return data, nil
}