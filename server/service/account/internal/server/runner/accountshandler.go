package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/record"
)

// GetAccountsHandler -
func (rnr *Runner) GetAccountsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get accounts handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Account
	var err error

	id := pp.ByName("account_id")

	// single resource
	if id != "" {

		l.Info("Getting account record ID >%s<", id)

		rec, err := m.(*model.Model).GetAccountRec(id, false)
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

		l.Info("Querying account records")

		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetAccountRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.AccountData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToAccountResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.AccountResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostAccountsHandler -
func (rnr *Runner) PostAccountsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post accounts handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("account_id")

	req := schema.AccountRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Account{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.AccountRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreateAccountRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToAccountResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.AccountResponse{
		Data: []schema.AccountData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutAccountsHandler -
func (rnr *Runner) PutAccountsHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put accounts handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("account_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetAccountRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.AccountRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.AccountRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdateAccountRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToAccountResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.AccountResponse{
		Data: []schema.AccountData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// AccountRequestDataToRecord -
func (rnr *Runner) AccountRequestDataToRecord(data schema.AccountData, rec *record.Account) error {

	rec.Name = data.Name
	rec.Email = data.Email
	rec.Provider = data.Provider
	rec.ProviderAccountID = data.ProviderAccountID

	return nil
}

// RecordToAccountResponseData -
func (rnr *Runner) RecordToAccountResponseData(accountRec *record.Account) (schema.AccountData, error) {

	data := schema.AccountData{
		ID:                accountRec.ID,
		Name:              accountRec.Name,
		Email:             accountRec.Email,
		Provider:          accountRec.Provider,
		ProviderAccountID: accountRec.ProviderAccountID,
		CreatedAt:         accountRec.CreatedAt,
		UpdatedAt:         accountRec.UpdatedAt.Time,
	}

	return data, nil
}
