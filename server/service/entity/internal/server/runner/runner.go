package runner

import (
	"net/http"

	"gitlab.com/alienspaces/holyragingmages/server/constant"

	"gitlab.com/alienspaces/holyragingmages/server/core/auth"
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
		// 0 - Get any, Administrator role or Default role, account ID not required
		{
			Method:      http.MethodGet,
			Path:        "/api/entities",
			HandlerFunc: r.GetEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				ValidateQueryParams: []string{
					"entity_type",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query entities.",
			},
		},
		// 1 - Get with entity ID, Administrator role or Default role, account ID not required
		{
			Method:      http.MethodGet,
			Path:        "/api/entities/:entity_id",
			HandlerFunc: r.GetEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an entity.",
			},
		},
		// 2 - Get any, Default or Administrator role, accountID required, account ID in path must match identity
		{
			Method:      http.MethodGet,
			Path:        "/api/accounts/:account_id/entities",
			HandlerFunc: r.GetAccountEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				ValidatePathParams: map[string]server.ValidatePathParam{
					"account_id": server.ValidatePathParam{
						MatchIdentity: true,
					},
				},
				ValidateQueryParams: []string{
					"entity_type",
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Query entities.",
			},
		},
		// 3 - Get with entity ID, Default or Administrator role, accountID required, account ID in path must match identity
		{
			Method:      http.MethodGet,
			Path:        "/api/accounts/:account_id/entities/:entity_id",
			HandlerFunc: r.GetAccountEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				ValidatePathParams: map[string]server.ValidatePathParam{
					"account_id": server.ValidatePathParam{
						MatchIdentity: true,
					},
				},
			},
			DocumentationConfig: server.DocumentationConfig{
				Document:    true,
				Description: "Get an entity.",
			},
		},
		// 4 - Post without entity ID, Default or Administrator role, accountID required, account ID in path must match identity
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts/:account_id/entities",
			HandlerFunc: r.PostAccountEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				ValidatePathParams: map[string]server.ValidatePathParam{
					"account_id": server.ValidatePathParam{
						MatchIdentity: true,
					},
				},
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
		// 5 - Post with entity ID, Default or Administrator role, accountID required, account ID in path must match identity
		{
			Method:      http.MethodPost,
			Path:        "/api/accounts/:account_id/entities/:entity_id",
			HandlerFunc: r.PostAccountEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				ValidatePathParams: map[string]server.ValidatePathParam{
					"account_id": server.ValidatePathParam{
						MatchIdentity: true,
					},
				},
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
		// 6 - Put with entity ID, Default or Administrator role, accountID required, account ID in path must match identity
		{
			Method:      http.MethodPut,
			Path:        "/api/accounts/:account_id/entities/:entity_id",
			HandlerFunc: r.PutAccountEntitiesHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				AuthTypes: []string{
					auth.AuthTypeJWT,
				},
				AuthRequireAnyRole: []string{
					constant.AuthRoleAdministrator,
					constant.AuthRoleDefault,
				},
				AuthRequireAllIdentities: []string{
					constant.AuthIdentityAccountID,
				},
				ValidatePathParams: map[string]server.ValidatePathParam{
					"account_id": server.ValidatePathParam{
						MatchIdentity: true,
					},
				},
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
		// 6 - Documentation
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
