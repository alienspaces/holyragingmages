package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/record"
)

// GetItemsHandler -
func (rnr *Runner) GetItemsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get items handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Item
	var err error

	id := pp.ByName("item_id")

	// single resource
	if id != "" {

		l.Info("Getting item record ID >%s<", id)

		rec, err := m.(*model.Model).GetItemRec(id, false)
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

		l.Info("Querying item records")

		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetItemRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.ItemData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToItemResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.ItemResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostItemsHandler -
func (rnr *Runner) PostItemsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post items handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("item_id")

	req := schema.ItemRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Item{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.ItemRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateItemRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToItemResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.ItemResponse{
		Data: []schema.ItemData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutItemsHandler -
func (rnr *Runner) PutItemsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put items handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("item_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetItemRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.ItemRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.ItemRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateItemRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToItemResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.ItemResponse{
		Data: []schema.ItemData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// ItemRequestDataToRecord -
func (rnr *Runner) ItemRequestDataToRecord(data schema.ItemData, rec *record.Item) error {

	rec.Name = data.Name
	rec.Description = data.Name

	return nil
}

// RecordToItemResponseData -
func (rnr *Runner) RecordToItemResponseData(itemRec *record.Item) (schema.ItemData, error) {

	data := schema.ItemData{
		ID:          itemRec.ID,
		Name:        itemRec.Name,
		Description: itemRec.Description,
		CreatedAt:   itemRec.CreatedAt,
		UpdatedAt:   itemRec.UpdatedAt.Time,
	}

	return data, nil
}
