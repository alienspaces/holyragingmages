package runner

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/record"
)

// MageResponse -
type MageResponse struct {
	server.Response
	Data []MageData `json:"data"`
}

// MageRequest -
type MageRequest struct {
	server.Request
	Data MageData `json:"data"`
}

// MageData -
type MageData struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name"`
	Strength     int       `json:"strength"`
	Dexterity    int       `json:"dexterity"`
	Intelligence int       `json:"intelligence"`
	Experience   int64     `json:"experience"`
	Coin         int64     `json:"coin"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

// GetMagesHandler -
func (rnr *Runner) GetMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get mages handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Mage
	var err error

	id := p.ByName("mage_id")

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
		q := r.URL.Query()

		params := make(map[string]interface{})
		for paramName, paramValue := range q {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetMageRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []MageData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToMageResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := MageResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostMagesHandler -
func (rnr *Runner) PostMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post mages handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("mage_id")

	req := MageRequest{}

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
	res := MageResponse{
		Data: []MageData{
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
func (rnr *Runner) PutMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put mages handler ** p >%#v< m >#%v<", p, m)

	id := p.ByName("mage_id")

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

	req := MageRequest{}

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
	res := MageResponse{
		Data: []MageData{
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
func (rnr *Runner) MageRequestDataToRecord(data MageData, rec *record.Mage) error {

	rec.Name = data.Name

	return nil
}

// RecordToMageResponseData -
func (rnr *Runner) RecordToMageResponseData(rec *record.Mage) (MageData, error) {

	data := MageData{
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
