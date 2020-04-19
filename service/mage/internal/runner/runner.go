package runner

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/payload"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/record"
)

// Runner -
type Runner struct {
	service.Runner
}

// Response -
type Response struct {
	service.Response
	Data []Data `json:"data"`
}

// Request -
type Request struct {
	service.Request
	Data Data `json:"data"`
}

// Data -
type Data struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Strength     int       `json:"strength"`
	Dexterity    int       `json:"dexterity"`
	Intelligence int       `json:"intelligence"`
	Experience   int64     `json:"experience"`
	Coin         int64     `json:"coin"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

// Fault -
type Fault struct {
	Error error
}

// ensure we comply with the Runnerer interface
var _ runnable.Runnable = &Runner{}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.RouterFunc = r.Router
	r.MiddlewareFunc = r.Middleware
	r.HandlerFunc = r.Handler
	r.ModellerFunc = r.Modeller
	r.PayloaderFunc = r.Payloader

	r.HandlerConfig = []service.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/api/mages",
			HandlerFunc:      r.GetMagesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/mages/:id",
			HandlerFunc:      r.GetMagesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/mages",
			HandlerFunc: r.PostMagesHandler,
			MiddlewareConfig: service.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/mages/:id",
			HandlerFunc: r.PutMagesHandler,
			MiddlewareConfig: service.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
	}

	return &r
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Mage Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h service.Handle) (service.Handle, error) {

	rnr.Log.Info("** Mage Middleware **")

	return h, nil
}

// Modeller -
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Mage Model **")

	m, err := model.NewModel(rnr.Config, rnr.Log, rnr.Store)
	if err != nil {
		rnr.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// Payloader -
func (rnr *Runner) Payloader() (payloader.Payloader, error) {

	rnr.Log.Info("** Payloader **")

	p, err := payload.NewPayload()
	if err != nil {
		rnr.Log.Warn("Failed new payloader >%v<", err)
		return nil, err
	}

	return p, nil
}

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Mage handler **")

	fmt.Fprint(w, "Hello from mage!\n")
}

// GetMagesHandler -
func (rnr *Runner) GetMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get mages handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Mage
	var err error

	id := p.ByName("id")

	// single resource
	if id != "" {

		rnr.Log.Info("Fetching resource ID >%s<", id)

		rec, err := m.(*model.Model).GetMageRec(id, false)
		if err != nil {
			rnr.Log.Warn("Failed getting mage record >%v<", err)

			// model error
			res := rnr.ErrorResponse(err)

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}

		// resource not found
		if rec == nil {
			rnr.Log.Warn("Get mage rec nil")

			// not found error
			res := rnr.ErrorNotFound(fmt.Errorf("Resource with ID >%s< not found", id))

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}

		recs = append(recs, rec)

	} else {

		rnr.Log.Info("Fetching all resources")

		recs, err = m.(*model.Model).GetMageRecs(nil, nil, false)
		if err != nil {
			rnr.Log.Warn("Failed getting mage records >%v<", err)

			// model error
			res := rnr.ErrorResponse(err)

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}
	}

	// assign response properties
	data := []Data{}
	for _, rec := range recs {
		data = append(data, Data{
			ID:           rec.ID,
			Name:         rec.Name,
			Strength:     rec.Strength,
			Dexterity:    rec.Dexterity,
			Intelligence: rec.Intelligence,
			Experience:   rec.Experience,
			Coin:         rec.Coin,
			CreatedAt:    rec.CreatedAt,
			UpdatedAt:    rec.UpdatedAt.Time,
		})
	}

	res := Response{
		Data: data,
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostMagesHandler -
func (rnr *Runner) PostMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post mages handler ** p >%#v< m >#%v<", p, m)

	req := Request{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)

		// system error
		res := rnr.ErrorSystem(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	rec := record.Mage{}

	// assign request properties
	rec.ID = req.Data.ID
	rec.Name = req.Data.Name

	err = m.(*model.Model).CreateMageRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed creating mage record >%v<", err)

		// model error
		res := rnr.ErrorResponse(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: []Data{
			{
				ID:           rec.ID,
				Name:         rec.Name,
				Strength:     rec.Strength,
				Dexterity:    rec.Dexterity,
				Intelligence: rec.Intelligence,
				Experience:   rec.Experience,
				Coin:         rec.Coin,
				CreatedAt:    rec.CreatedAt,
				UpdatedAt:    rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutMagesHandler -
func (rnr *Runner) PutMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Put mages handler ** p >%#v< m >#%v<", p, m)

	id := p.ByName("id")

	rnr.Log.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetMageRec(id, false)
	if err != nil {
		rnr.Log.Warn("Failed getting mage record >%v<", err)

		// model error
		res := rnr.ErrorResponse(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// resource not found
	if rec == nil {
		rnr.Log.Warn("Get mage rec nil")

		// not found error
		res := rnr.ErrorNotFound(fmt.Errorf("Resource with ID >%s< not found", id))

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	req := Request{}

	err = rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)

		// system error
		res := rnr.ErrorSystem(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// assign request properties
	rec.Name = req.Data.Name

	err = m.(*model.Model).UpdateMageRec(rec)
	if err != nil {
		rnr.Log.Warn("Failed updating mage record >%v<", err)

		// model error
		res := rnr.ErrorResponse(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: []Data{
			{
				ID:           rec.ID,
				Name:         rec.Name,
				Strength:     rec.Strength,
				Dexterity:    rec.Dexterity,
				Intelligence: rec.Intelligence,
				Experience:   rec.Experience,
				Coin:         rec.Coin,
				CreatedAt:    rec.CreatedAt,
				UpdatedAt:    rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}
