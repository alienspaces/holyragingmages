package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/payload"
	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
)

// Runner -
type Runner struct {
	server.Runner
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

	r.HandlerConfig = []server.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/api/templates",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/templates/:template_id",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/templates",
			HandlerFunc: r.PostTemplatesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/templates/:template_id",
			HandlerFunc: r.PostTemplatesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/templates/:template_id",
			HandlerFunc: r.PutTemplatesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
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
func (rnr *Runner) Middleware(h server.Handle) (server.Handle, error) {

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