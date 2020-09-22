package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
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

	r.HandlerConfig = []server.HandlerConfig{
		// Authentication
		{
			Method:      http.MethodGet,
			Path:        "/api/auth",
			HandlerFunc: r.PostAuthHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "auth",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Authenticate OAuth provider token.",
			},
		},
		// Accounts
		{
			Method:           http.MethodGet,
			Path:             "/api/accounts",
			HandlerFunc:      r.GetAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query accounts.",
			},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/accounts/:account_id",
			HandlerFunc:      r.GetAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an account.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts",
			HandlerFunc: r.PostAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a account.",
			},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts/:account_id",
			HandlerFunc: r.PostAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a account.",
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/accounts/:account_id",
			HandlerFunc: r.PutAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update a account.",
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

	l.Info("** Account Model **")

	m, err := model.NewModel(rnr.Config, l, rnr.Store)
	if err != nil {
		l.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}
