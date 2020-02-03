package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
)

// Runner -
type Runner struct {
	service.Runner
}

// ensure we comply with the Runnerer interface
var _ service.Runnable = &Runner{}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.RouterFunc = r.Router
	r.MiddlewareFunc = r.Middleware
	r.HandlerFunc = r.Handler
	r.ModelFunc = r.Model

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

// Model -
func (rnr *Runner) Model(c service.Configurer, l service.Logger, s service.Storer) (service.Modeller, error) {

	rnr.Log.Info("** Template Model **")

	m, err := model.NewModel(c, l, s)
	if err != nil {
		rnr.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m service.Modeller) {

	rnr.Log.Info("** Template handler **")

	fmt.Fprint(w, "Hello from template!\n")
}

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m service.Modeller) {

	rnr.Log.Info("** Get templates handler ** p >%#v< m >%#v<", p, m)

	fmt.Fprint(w, "Hello from GET templates handler!\n", p)
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m service.Modeller) {

	data := r.Context().Value(service.ContextKeyData)

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v< data >%v<", p, m, data)

	fmt.Fprint(w, "Hello from Post templates handler!\n", p, "\n", data)
}

// PutTemplatesHandler -
func (rnr *Runner) PutTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m service.Modeller) {

	data := r.Context().Value(service.ContextKeyData)

	rnr.Log.Info("** Put templates handler ** p >%#v< m >#%v< data >%v<", p, m, data)

	fmt.Fprint(w, "Hello from Put templates handler!\n", p, "\n", data)
}
