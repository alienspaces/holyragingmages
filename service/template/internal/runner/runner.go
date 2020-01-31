package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/service"
)

// Runner -
type Runner struct {
	service.Runner
}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.RouterFunc = r.Router
	r.MiddlewareFunc = r.Middleware
	r.HandlerFunc = r.Handler

	r.HandlerConfig = []service.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/templates",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:           http.MethodGet,
			Path:             "/templates/:id",
			HandlerFunc:      r.GetTemplatesHandler,
			MiddlewareConfig: service.MiddlewareConfig{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/templates",
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
			Path:        "/templates/:id",
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
func (rnr *Runner) Middleware(h httprouter.Handle) (httprouter.Handle, error) {

	rnr.Log.Info("** Template Middleware **")

	return h, nil
}

// Handler -
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	rnr.Log.Info("** Template handler **")

	fmt.Fprint(w, "Hello from template!\n")
}

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	rnr.Log.Info("** Get templates handler ** params >%v<", params)

	fmt.Fprint(w, "Hello from GET templates handler!\n", params)
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	data := r.Context().Value(service.ContextKeyData)

	rnr.Log.Info("** Post templates handler ** params >%v< data >%v<", params, data)

	fmt.Fprint(w, "Hello from Post templates handler!\n", params, "\n", data)
}

// PutTemplatesHandler -
func (rnr *Runner) PutTemplatesHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	data := r.Context().Value(service.ContextKeyData)

	rnr.Log.Info("** Put templates handler ** params >%v< data >%v<", params, data)

	fmt.Fprint(w, "Hello from Put templates handler!\n", params, "\n", data)
}
