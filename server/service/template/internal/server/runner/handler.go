package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/template/internal/record"
)

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get templates handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Template
	var err error

	id := pp.ByName("template_id")

	// single resource
	if id != "" {

		l.Info("Getting template record ID >%s<", id)

		rec, err := m.(*model.Model).GetTemplateRec(id, false)
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

		l.Info("Querying template records")

		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetTemplateRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.TemplateData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToTemplateResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.TemplateResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post templates handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("template_id")

	req := schema.TemplateRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Template{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.TemplateRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateTemplateRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToTemplateResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.TemplateResponse{
		Data: []schema.TemplateData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutTemplatesHandler -
func (rnr *Runner) PutTemplatesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put templates handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("template_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetTemplateRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.TemplateRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.TemplateRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateTemplateRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToTemplateResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.TemplateResponse{
		Data: []schema.TemplateData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// TemplateRequestDataToRecord -
func (rnr *Runner) TemplateRequestDataToRecord(data schema.TemplateData, rec *record.Template) error {

	return nil
}

// RecordToTemplateResponseData -
func (rnr *Runner) RecordToTemplateResponseData(templateRec *record.Template) (schema.TemplateData, error) {

	data := schema.TemplateData{
		ID:        templateRec.ID,
		CreatedAt: templateRec.CreatedAt,
		UpdatedAt: templateRec.UpdatedAt.Time,
	}

	return data, nil
}
