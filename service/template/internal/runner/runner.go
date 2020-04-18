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
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
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
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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
			Path:             "/api/templates",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/templates/:id",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/templates",
			HandlerFunc: r.PostTemplatesHandler,
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
			Path:        "/api/templates/:id",
			HandlerFunc: r.PutTemplatesHandler,
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

	rnr.Log.Info("** Template Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h service.Handle) (service.Handle, error) {

	rnr.Log.Info("** Template Middleware **")

	return h, nil
}

// Modeller -
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Template Model **")

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

	rnr.Log.Info("** Template handler **")

	fmt.Fprint(w, "Hello from template!\n")
}

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get templates handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Template
	var err error

	id := p.ByName("id")

	// single resource
	if id != "" {

		rnr.Log.Info("Fetching resource ID >%s<", id)

		rec, err := m.(*model.Model).GetTemplateRec(id, false)
		if err != nil {
			rnr.Log.Warn("Get template rec err >%v<", err)

			// system error
			res := rnr.ErrorSystem(err)

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}

		// resource not found
		if rec == nil {
			rnr.Log.Warn("Get template rec nil")

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

		recs, err = m.(*model.Model).GetTemplateRecs(nil, nil, false)
		if err != nil {

			// system error
			res := rnr.ErrorSystem(err)

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
			ID:        rec.ID,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt.Time,
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

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v<", p, m)

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

	rec := record.Template{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).CreateTemplateRec(&rec)
	if err != nil {
		// system error
		res := rnr.ErrorSystem(err)

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
				ID:        rec.ID,
				CreatedAt: rec.CreatedAt,
				UpdatedAt: rec.UpdatedAt.Time,
			},
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

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v<", p, m)

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

	rec := record.Template{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).UpdateTemplateRec(&rec)
	if err != nil {
		// system error
		res := rnr.ErrorSystem(err)

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
				ID:        rec.ID,
				CreatedAt: rec.CreatedAt,
				UpdatedAt: rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}
