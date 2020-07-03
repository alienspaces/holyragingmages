package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/record"
)

// GetMagesHandler -
func (rnr *Runner) GetMagesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get mages handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Mage
	var err error

	id := pp.ByName("mage_id")

	// single resource
	if id != "" {

		l.Info("Getting mage record ID >%s<", id)

		rec, err := m.(*model.Model).GetMageRec(id, false)
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

		recs, err = m.(*model.Model).GetMageRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.MageData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToMageResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.MageResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostMagesHandler -
func (rnr *Runner) PostMagesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post mages handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("mage_id")

	req := schema.MageRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Mage{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.MageRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateMageRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToMageResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.MageResponse{
		Data: []schema.MageData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutMagesHandler -
func (rnr *Runner) PutMagesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put mages handler ** p >%#v< m >#%v<", pp, m)

	id := pp.ByName("mage_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetMageRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.MageRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.MageRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateMageRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToMageResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.MageResponse{
		Data: []schema.MageData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// MageRequestDataToRecord -
func (rnr *Runner) MageRequestDataToRecord(data schema.MageData, rec *record.Mage) error {

	rec.Name = data.Name

	return nil
}

// RecordToMageResponseData -
func (rnr *Runner) RecordToMageResponseData(rec *record.Mage) (schema.MageData, error) {

	data := schema.MageData{
		ID:           rec.ID,
		Name:         rec.Name,
		Strength:     rec.Strength,
		Dexterity:    rec.Dexterity,
		Intelligence: rec.Intelligence,
		Experience:   rec.Experience,
		Coin:         rec.Coin,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt.Time,
	}

	return data, nil
}
