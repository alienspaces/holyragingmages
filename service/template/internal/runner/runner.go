package runner

import (
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
	ID string `json:"id"`
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
