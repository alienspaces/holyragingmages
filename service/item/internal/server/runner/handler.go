package runner

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/item/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/item/internal/record"
)

// ItemResponse -
type ItemResponse struct {
	server.Response
	Data []ItemData `json:"data"`
}

// ItemRequest -
type ItemRequest struct {
	server.Request
	Data ItemData `json:"data"`
}

// ItemData -
type ItemData struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// GetItemsHandler -
func (rnr *Runner) GetItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get items handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Item
	var err error

	id := p.ByName("item_id")

	// single resource
	if id != "" {

		rnr.Log.Info("Getting item ID >%s<", id)

		rec, err := m.(*model.Model).GetItemRec(id, false)
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

		rnr.Log.Info("Getting all item records")

		// query parameters
		q := r.URL.Query()

		params := make(map[string]interface{})
		for paramName, paramValue := range q {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetItemRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(w, err)
			return
		}
	}

	// assign response properties
	data := []ItemData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToItemResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(w, err)
			return
		}

		data = append(data, responseData)
	}

	res := ItemResponse{
		Data: data,
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostItemsHandler -
func (rnr *Runner) PostItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post items handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("item_id")

	req := ItemRequest{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	rec := record.Item{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.ItemRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).CreateItemRec(&rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToItemResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := ItemResponse{
		Data: []ItemData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutItemsHandler -
func (rnr *Runner) PutItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Put items handler ** p >%#v< m >#%v<", p, m)

	// parameters
	id := p.ByName("item_id")

	rnr.Log.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetItemRec(id, false)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(w, id)
		return
	}

	req := ItemRequest{}

	err = rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// record data
	err = rnr.ItemRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	err = m.(*model.Model).UpdateItemRec(rec)
	if err != nil {
		rnr.WriteModelError(w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToItemResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(w, err)
		return
	}

	// assign response properties
	res := ItemResponse{
		Data: []ItemData{
			responseData,
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// ItemRequestDataToRecord -
func (rnr *Runner) ItemRequestDataToRecord(data ItemData, rec *record.Item) error {

	rec.Name = data.Name
	rec.Description = data.Name

	return nil
}

// RecordToItemResponseData -
func (rnr *Runner) RecordToItemResponseData(itemRec *record.Item) (ItemData, error) {

	data := ItemData{
		ID:          itemRec.ID,
		Name:        itemRec.Name,
		Description: itemRec.Description,
		CreatedAt:   itemRec.CreatedAt,
		UpdatedAt:   itemRec.UpdatedAt.Time,
	}

	return data, nil
}