package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/core/prepare"
	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/service/item/internal/model"
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
	r.PreparerFunc = r.Preparer
	r.ModellerFunc = r.Modeller
	r.PayloaderFunc = r.Payloader

	r.HandlerConfig = []server.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/api/items",
			QueryParams:      []string{"name"},
			HandlerFunc:      r.GetItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query items.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/items/:item_id",
			HandlerFunc:      r.GetItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an item.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/items",
			HandlerFunc: r.PostItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create an item.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/items/:item_id",
			HandlerFunc: r.PostItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create an item.",
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/items/:item_id",
			HandlerFunc: r.PutItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update an item.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api",
			HandlerFunc:      r.GetDocumentationHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
		},
	}

	return &r
}

// Preparer -
func (rnr *Runner) Preparer(l logger.Logger) (preparer.Preparer, error) {

	l.Info("** Item Preparer **")

	p, err := prepare.NewPrepare(l)
	if err != nil {
		l.Warn("Failed new preparer >%v<", err)
		return nil, err
	}

	return p, nil
}

// Modeller -
func (rnr *Runner) Modeller(l logger.Logger) (modeller.Modeller, error) {

	l.Info("** Item Model **")

	m, err := model.NewModel(rnr.Config, l, rnr.Store)
	if err != nil {
		l.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}
