package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/core/prepare"
	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/preparer"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/service/spell/internal/model"
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
			Path:             "/api/spells",
			HandlerFunc:      r.GetSpellsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query spells.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/spells/:spell_id",
			HandlerFunc:      r.GetSpellsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get a spell.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/spells",
			HandlerFunc: r.PostSpellsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a spell.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/spells/:spell_id",
			HandlerFunc: r.PostSpellsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a spell.",
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/spells/:spell_id",
			HandlerFunc: r.PutSpellsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update a spell.",
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

	l.Info("** Spell Preparer **")

	p, err := prepare.NewPrepare(l)
	if err != nil {
		l.Warn("Failed new preparer >%v<", err)
		return nil, err
	}

	return p, nil
}

// Modeller -
func (rnr *Runner) Modeller(l logger.Logger) (modeller.Modeller, error) {

	l.Info("** Spell Model **")

	m, err := model.NewModel(rnr.Config, l, rnr.Store)
	if err != nil {
		l.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}
