package runner

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// TemplateResponse -
type TemplateResponse struct {
	server.Response
	Data []TemplateData `json:"data"`
}

// TemplateRequest -
type TemplateRequest struct {
	server.Request
	Data TemplateData `json:"data"`
}

// TemplateData -
type TemplateData struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get templates handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Template
	var err error

	id := p.ByName("template_id")

	// single resource
	if id != "" {

		rnr.Log.Info("Getting template ID >%s<", id)

		rec, err := m.(*model.Model).GetTemplateRec(id, false)
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

		rnr.Log.Info("Getting all template records")

		recs, err = m.(*model.Model).GetTemplateRecs(nil, nil, false)
		if err != nil {
			rnr.WriteModelError(w, err)
			return
		}
	}

	// assign response properties
	data := []TemplateData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToTemplateResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(w, err)
			return
		}

		data = append(data, responseData)
	}

	res := TemplateResponse{
		Data: data,
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("template_id")

	req := TemplateRequest{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	rec := record.Template{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.TemplateRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).CreateTemplateRec(&rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToTemplateResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := TemplateResponse{
		Data: []TemplateData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutTemplatesHandler -
func (rnr *Runner) PutTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Put templates handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("template_id")

	rnr.Log.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetTemplateRec(id, false)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(w, id)
		return
	}

	req := TemplateRequest{}

	err = rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// record data
	err = rnr.TemplateRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).UpdateTemplateRec(rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToTemplateResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := TemplateResponse{
		Data: []TemplateData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// TemplateRequestDataToRecord -
func (rnr *Runner) TemplateRequestDataToRecord(data TemplateData, rec *record.Template) error {

	return nil
}

// RecordToTemplateResponseData -
func (rnr *Runner) RecordToTemplateResponseData(templateRec *record.Template) (TemplateData, error) {

	data := TemplateData{
		ID:        templateRec.ID,
		CreatedAt: templateRec.CreatedAt,
		UpdatedAt: templateRec.UpdatedAt.Time,
	}

	return data, nil
}