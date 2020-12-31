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

// GetEntitiesHandler - Admininstrator role only, account ID not required
func (rnr *Runner) GetEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get entities handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Entity
	var err error

	id := pp.ByName("entity_id")

	// single resource
	if id != "" {

		l.Info("Getting entity record ID >%s<", id)

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

		l.Info("Querying entity records")

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

// GetAccountEntitiesHandler - Default or Administrator role, accountID required, account ID in path must match identity
func (rnr *Runner) GetAccountEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get entities handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Entity
	var err error

	accountID := pp.ByName("account_id")
	entityID := pp.ByName("entity_id")

	// single resource
	if accountID != "" && entityID != "" {

		l.Info("Getting account ID >%s< entity ID >%s< record ", accountID, entityID)

		rec, err := m.(*model.Model).GetEntityRec(entityID, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		// resource not found
		if rec == nil {
			l.Warn("Record entity ID >%s< not found", entityID)
			rnr.WriteNotFoundError(l, w, entityID)
			return
		}

		l.Info("Checking account ID >%s< of entity ID >%s< record ", accountID, entityID)
		if rec.AccountID != accountID {
			l.Warn("Record entity ID >%s< with account ID >%s< does not match requested account ID >%s<", entityID, rec.AccountID, accountID)
			rnr.WriteNotFoundError(l, w, entityID)
			return
		}

		// entity belongs to account
		recs = append(recs, rec)

	} else {

		l.Info("Querying entity records")

		// query parameters
		params := map[string]interface{}{
			"account_id": accountID,
		}
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
			l.Info("Querying entity records with param name >%s< value >%v<", paramName, paramValue)
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

// PostAccountEntitiesHandler -
func (rnr *Runner) PostAccountEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post entities handler ** p >%#v< m >#%v<", pp, m)

	// Path parameters
	entityID := pp.ByName("entity_id")
	accountID := pp.ByName("account_id")

	req := schema.EntityRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Entity{}

	// Assign request properties
	rec.ID = entityID
	rec.AccountID = accountID

	// Record data
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

	// Response data
	responseData, err := rnr.RecordToEntityResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// Assign response properties
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

// PutAccountEntitiesHandler -
func (rnr *Runner) PutAccountEntitiesHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put entities handler ** p >%#v< m >#%v<", pp, m)

	entityID := pp.ByName("entity_id")
	accountID := pp.ByName("account_id")

	l.Info("Updating resource ID >%s<", entityID)

	rec, err := m.(*model.Model).GetEntityRec(entityID, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// Resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, entityID)
		return
	}

	// Record account ID must match path paramter
	l.Info("Checking account ID >%s< of entity ID >%s< record ", accountID, entityID)

	if rec.AccountID != accountID {
		l.Warn("Record entity ID >%s< with account ID >%s< does not match requested account ID >%s<", entityID, rec.AccountID, accountID)
		rnr.WriteNotFoundError(l, w, entityID)
		return
	}

	req := schema.EntityRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// Record data
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

	// Response data
	responseData, err := rnr.RecordToEntityResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// Assign response properties
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

	// NOTE: AccountID is sourced from path parameters

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
		AccountID:        rec.AccountID,
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
