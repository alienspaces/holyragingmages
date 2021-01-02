package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/constant"
	"gitlab.com/alienspaces/holyragingmages/server/core/auth"
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
		// 0 - Authentication
		{
			Method:      http.MethodPost,
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
		// 1 - Refresh Authentication
		{
			Method:      http.MethodPost,
			Path:        "/api/auth-refresh",
			HandlerFunc: r.PostAuthRefreshHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "authrefresh",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Refresh authentication token.",
			},
		},

		// TODO: Provide different default and administrator account handlers constraining record access as appropriate

		// 2 - Accounts - Get many
		{
			Method:      http.MethodGet,
			Path:        "/api/accounts",
			HandlerFunc: r.GetAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleDefault,
					constant.AuthRoleAdministrator,
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query accounts.",
			},
		},
		// 3 - Accounts - Get one
		{
			Method:      http.MethodGet,
			Path:        "/api/accounts/:account_id",
			HandlerFunc: r.GetAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleDefault,
					constant.AuthRoleAdministrator,
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an account.",
			},
		},
		// 4 - Accounts - Create one without ID
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts",
			HandlerFunc: r.PostAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleDefault,
					constant.AuthRoleAdministrator,
				},
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "account-main.schema.json",
				ValidateSchemaReferences: []string{
					"account-data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a account.",
			},
		},
		// 5 - Accounts - Create one with ID
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts/:account_id",
			HandlerFunc: r.PostAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleDefault,
					constant.AuthRoleAdministrator,
				},
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "account-main.schema.json",
				ValidateSchemaReferences: []string{
					"account-data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Create a account.",
			},
		},
		// 6 - Accounts - Update one
		{
			Method:      http.MethodPut,
			Path:        "/api/accounts/:account_id",
			HandlerFunc: r.PutAccountsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleDefault,
					constant.AuthRoleAdministrator,
				},
				ValidateSchemaLocation: "account",
				ValidateSchemaMain:     "account-main.schema.json",
				ValidateSchemaReferences: []string{
					"account-data.schema.json",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Update a account.",
			},
		},
		// 7 - Documentation
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
