package server

import (
	"context"
)

// All context keys are of type string
type contextKey string

// Context keys
const (
	contextRolesKey    contextKey = "authRoles"
	contextIdentityKey contextKey = "authIdentity"
)

// addRolesContext -
func (rnr *Runner) addRolesContext(ctx context.Context, roles []string) (context.Context, error) {

	ctx = context.WithValue(ctx, contextRolesKey, roles)

	return ctx, nil
}

// getRolesContext -
func (rnr *Runner) getRolesContext(ctx context.Context) ([]string, error) {

	roles := ctx.Value(contextRolesKey)
	if roles != nil {
		return roles.([]string), nil
	}

	return nil, nil
}

// addIdentityContext -
func (rnr *Runner) addIdentityContext(ctx context.Context, identity map[string]interface{}) (context.Context, error) {

	ctx = context.WithValue(ctx, contextIdentityKey, identity)

	return ctx, nil
}

// getIdentityContext -
func (rnr *Runner) getIdentityContext(ctx context.Context) (map[string]interface{}, error) {

	identity := ctx.Value(contextRolesKey)
	if identity != nil {
		return identity.(map[string]interface{}), nil
	}

	return nil, nil
}
