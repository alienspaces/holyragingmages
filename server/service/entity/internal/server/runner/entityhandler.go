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

	var entityRecs []*record.Entity

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

		entityRecs = append(entityRecs, rec)

	} else {

		l.Info("Querying entity records")

		// query parameters
		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		entityRecs, err = m.(*model.Model).GetEntityRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.EntityData{}
	for _, entityRec := range entityRecs {

		// Get associated account entity record if it exists
		accountEntityRecs, err := m.(*model.Model).GetAccountEntityRecs(
			map[string]interface{}{
				"entity_id": entityRec.ID,
			},
			nil, false,
		)

		var accountEntityRec *record.AccountEntity
		if len(accountEntityRecs) == 1 {
			accountEntityRec = accountEntityRecs[0]
		}

		// response data
		responseData, err := rnr.RecordToEntityResponseData(accountEntityRec, entityRec)
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

	var entityRecs []*record.Entity
	var accountEntityRecs []*record.AccountEntity

	var err error

	accountID := pp.ByName("account_id")
	entityID := pp.ByName("entity_id")

	// single resource
	if accountID != "" && entityID != "" {

		l.Info("Getting account ID >%s< entity ID >%s< record ", accountID, entityID)

		accountEntityRecs, err := m.(*model.Model).GetAccountEntityRecs(
			map[string]interface{}{
				"account_id": accountID,
				"entity_id":  entityID,
			},
			nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		// resource not found
		if len(accountEntityRecs) == 0 {
			l.Warn("Record account ID >%s< entity ID >%s< not found", accountID, entityID)
			rnr.WriteNotFoundError(l, w, entityID)
			return
		}

		accountEntityRec := accountEntityRecs[0]

		entityRec, err := m.(*model.Model).GetEntityRec(entityID, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		// Entity records
		entityRecs = append(entityRecs, entityRec)

		// Account entity records
		accountEntityRecs = append(accountEntityRecs, accountEntityRec)

	} else {

		l.Info("Querying entity records")

		// query parameters
		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		queryEntityRecs, err := m.(*model.Model).GetEntityRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		for _, entityRec := range queryEntityRecs {

			// query parameters
			params := map[string]interface{}{
				"entity_id":  entityRec.ID,
				"account_id": accountID,
			}

			queryAccountEntityRecs, err := m.(*model.Model).GetAccountEntityRecs(params, nil, false)
			if err != nil {
				rnr.WriteModelError(l, w, err)
				return
			}

			if len(queryAccountEntityRecs) == 1 {
				// Entity records
				entityRecs = append(entityRecs, entityRec)

				// Account entity records
				accountEntityRecs = append(accountEntityRecs, queryAccountEntityRecs[0])
			}
		}
	}

	// assign response properties
	data := []schema.EntityData{}
	for idx, entityRec := range entityRecs {

		accountEntityRec := accountEntityRecs[idx]

		// response data
		responseData, err := rnr.RecordToEntityResponseData(accountEntityRec, entityRec)
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

	entityRec := record.Entity{}

	// Assign request properties
	entityRec.ID = entityID

	// Record data
	err = rnr.EntityRequestDataToRecord(req.Data, &entityRec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// Default entity type to player character
	if entityRec.EntityType == "" {
		l.Info("Defaulting entity type to >%s<", record.EntityTypePlayerCharacter)
		entityRec.EntityType = record.EntityTypePlayerCharacter
	}

	err = m.(*model.Model).CreateEntityRec(&entityRec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	accountEntityRec := record.AccountEntity{
		EntityID:  entityRec.ID,
		AccountID: accountID,
	}

	err = m.(*model.Model).CreateAccountEntityRec(&accountEntityRec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// Response data
	responseData, err := rnr.RecordToEntityResponseData(&accountEntityRec, &entityRec)
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

	l.Info("Updating resource account ID >%s< entity ID >%s<", accountID, entityID)

	// Account entity record
	l.Info("Getting account ID >%s< entity ID >%s< record ", accountID, entityID)

	accountEntityRecs, err := m.(*model.Model).GetAccountEntityRecs(
		map[string]interface{}{
			"account_id": accountID,
			"entity_id":  entityID,
		},
		nil, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if len(accountEntityRecs) == 0 {
		l.Warn("Record account ID >%s< entity ID >%s< not found", accountID, entityID)
		rnr.WriteNotFoundError(l, w, entityID)
		return
	}

	accountEntityRec := accountEntityRecs[0]

	entityRec, err := m.(*model.Model).GetEntityRec(entityID, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// Resource not found
	if entityRec == nil {
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
	err = rnr.EntityRequestDataToRecord(req.Data, entityRec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateEntityRec(entityRec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// Response data
	responseData, err := rnr.RecordToEntityResponseData(accountEntityRec, entityRec)
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
func (rnr *Runner) RecordToEntityResponseData(accountEntityRec *record.AccountEntity, entityRec *record.Entity) (schema.EntityData, error) {

	data := schema.EntityData{
		ID:               entityRec.ID,
		AccountID:        accountEntityRec.AccountID,
		Name:             entityRec.Name,
		Strength:         entityRec.Strength,
		Dexterity:        entityRec.Dexterity,
		Intelligence:     entityRec.Intelligence,
		AttributePoints:  entityRec.AttributePoints,
		ExperiencePoints: entityRec.ExperiencePoints,
		Coins:            entityRec.Coins,
		CreatedAt:        entityRec.CreatedAt,
		UpdatedAt:        entityRec.UpdatedAt.Time,
	}

	return data, nil
}
