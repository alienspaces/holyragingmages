package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/service/entity/internal/model"
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
	r.ModellerFunc = r.Modeller

	r.HandlerConfig = []server.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/api/entities",
			HandlerFunc:      r.GetEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query entities.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/entities/:entity_id",
			HandlerFunc:      r.GetEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an entity.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/entities",
			HandlerFunc: r.PostEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "entity",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create an entity.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/entities/:entity_id",
			HandlerFunc: r.PostEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "entity",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create an entity.",
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/entities/:entity_id",
			HandlerFunc: r.PutEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "entity",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update a entity.",
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

// Modeller -
func (rnr *Runner) Modeller(l logger.Logger) (modeller.Modeller, error) {

	l.Info("** Entity Model **")

	m, err := model.NewModel(rnr.Config, l, rnr.Store)
	if err != nil {
		l.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}
