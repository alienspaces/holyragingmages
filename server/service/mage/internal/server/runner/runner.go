package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/core/prepare"
	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/service/mage/internal/model"
)

// Runner -
type Runner struct {
	server.Runner
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
			Path:             "/api/mages",
			HandlerFunc:      r.GetMagesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query mages.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/mages/:mage_id",
			HandlerFunc:      r.GetMagesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get a mage.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/mages",
			HandlerFunc: r.PostMagesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema/mage",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a mage.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/mages/:mage_id",
			HandlerFunc: r.PostMagesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema/mage",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a mage.",
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/mages/:mage_id",
			HandlerFunc: r.PutMagesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema/mage",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update a mage.",
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

	l.Info("** Mage Preparer **")

	p, err := prepare.NewPrepare(l)
	if err != nil {
		l.Warn("Failed new preparer >%v<", err)
		return nil, err
	}

	return p, nil
}

// Modeller -
func (rnr *Runner) Modeller(l logger.Logger) (modeller.Modeller, error) {

	l.Info("** Mage Model **")

	m, err := model.NewModel(rnr.Config, l, rnr.Store)
	if err != nil {
		l.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}