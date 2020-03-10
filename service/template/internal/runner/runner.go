package runner

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/payload"
	"gitlab.com/alienspaces/holyragingmages/common/prepare"
	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/type/configurer"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/common/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/common/type/storer"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// Runner -
type Runner struct {
	service.Runner
}

// Response -
type Response struct {
	Data  Data `json:"data"`
	Error error
}

// CollectionResponse -
type CollectionResponse struct {
	Data []Data `json:"data"`
}

// Request -
type Request struct {
	Data Data `json:"data"`
}

// Data -
type Data struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
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
	r.PreparerFunc = r.Preparer
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

// Preparer -
func (rnr *Runner) Preparer(l logger.Logger, tx *sqlx.Tx) (preparer.Preparer, error) {

	rnr.Log.Info("** Template Model **")

	p, err := prepare.NewPrepare(l, tx)
	if err != nil {
		rnr.Log.Warn("Failed new preparer >%v<", err)
		return nil, err
	}

	return p, nil
}

// Modeller -
func (rnr *Runner) Modeller(c configurer.Configurer, l logger.Logger, s storer.Storer) (modeller.Modeller, error) {

	rnr.Log.Info("** Template Model **")

	m, err := model.NewModel(c, l, s)
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

	fmt.Fprint(w, "Hello from GET templates handler!\n", p)
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v<", p, m)

	req := Request{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)
		fmt.Fprint(w, "Failed reading request\n", err)
		return
	}

	rec := record.Template{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).CreateTemplateRec(&rec)
	if err != nil {
		res := Response{
			Error: err,
		}
		err = rnr.WriteResponse(w, &res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			fmt.Fprint(w, "Failed writing response\n", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: Data{
			ID:        rec.ID,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt.String,
		},
	}

	err = rnr.WriteResponse(w, &res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		fmt.Fprint(w, "Failed writing response\n", err)
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
		fmt.Fprint(w, "Failed reading request\n", err)
		return
	}

	rec := record.Template{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).UpdateTemplateRec(&rec)
	if err != nil {
		res := Response{
			Error: err,
		}
		err = rnr.WriteResponse(w, &res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			fmt.Fprint(w, "Failed writing response\n", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: Data{
			ID:        rec.ID,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt.String,
		},
	}

	err = rnr.WriteResponse(w, &res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		fmt.Fprint(w, "Failed writing response\n", err)
		return
	}
}
