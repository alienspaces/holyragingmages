package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/record"
)

// GetEntitiesHandler -
func (rnr *Runner) GetEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get mages handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Entity
	var err error

	id := pp.ByName("mage_id")

	// single resource
	if id != "" {

		l.Info("Getting mage record ID >%s<", id)

		rec, err := m.(*model.Model).GetEntityRec(id, false)
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

		l.Info("Querying mage records")

		// query parameters
		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetEntityRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.EntityData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToEntityResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.EntityResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostEntitiesHandler -
func (rnr *Runner) PostEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post mages handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("mage_id")

	req := schema.EntityRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Entity{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.EntityRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateEntityRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToEntityResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.EntityResponse{
		Data: []schema.EntityData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutEntitiesHandler -
func (rnr *Runner) PutEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put mages handler ** p >%#v< m >#%v<", pp, m)

	id := pp.ByName("mage_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetEntityRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.EntityRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.EntityRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateEntityRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToEntityResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.EntityResponse{
		Data: []schema.EntityData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// EntityRequestDataToRecord -
func (rnr *Runner) EntityRequestDataToRecord(data schema.EntityData, rec *record.Entity) error {

	rec.Name = data.Name
	rec.Strength = data.Strength
	rec.Dexterity = data.Dexterity
	rec.Intelligence = data.Intelligence

	return nil
}

// RecordToEntityResponseData -
func (rnr *Runner) RecordToEntityResponseData(rec *record.Entity) (schema.EntityData, error) {

	data := schema.EntityData{
		ID:               rec.ID,
		Name:             rec.Name,
		Strength:         rec.Strength,
		Dexterity:        rec.Dexterity,
		Intelligence:     rec.Intelligence,
		AttributePoints:  rec.AttributePoints,
		ExperiencePoints: rec.ExperiencePoints,
		Coins:            rec.Coins,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt.Time,
	}

	return data, nil
}
