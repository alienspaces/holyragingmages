package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// authzIdentityCache - path, method
var authzIdentityCache map[string]map[string][]string

// authzRoleCache - path, method
var authzRoleCache map[string]map[string][]string

// Authz -
func (rnr *Runner) Authz(hc HandlerConfig, h HandlerFunc) (HandlerFunc, error) {

	// Cache authz configuration
	err := rnr.authzCacheConfig(hc)
	if err != nil {
		rnr.Log.Warn("Failed caching authz config >%v<", err)
		return nil, err
	}

	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Debug("** Authz **")

		err := rnr.handleAuthz(r, l, m, hc)
		if err != nil {
			l.Warn("Failed authz >%v<", err)
			rnr.WriteUnauthorizedError(l, w, err)
			return
		}

		h(w, r, pp, qp, l, m)
	}

	return handle, nil
}

// handleAuthz -
func (rnr *Runner) handleAuthz(r *http.Request, l logger.Logger, m modeller.Modeller, hc HandlerConfig) error {

	// Authorization may need to check identities and roles from context
	ctx := r.Context()

	if authzIdentityCache == nil && authzRoleCache == nil {
		l.Info("Authz not configured")
		return nil
	}

	authzIdentityMethods := authzIdentityCache[hc.Path]
	if authzIdentityMethods != nil {
		authzIdentities := authzIdentityMethods[hc.Method]
		if authzIdentities != nil {
			for _, identityKey := range authzIdentities {
				identityValue, err := rnr.getContextIdentityValue(ctx, identityKey)
				if err != nil {
					l.Warn("Failed checking context identity key >%s< >%v<", identityKey, err)
					return err
				}
				if identityValue == nil {
					msg := fmt.Sprintf("Context missing identity key >%s< value", identityKey)
					l.Warn(msg)
					return fmt.Errorf(msg)
				}
				l.Info("Have identity key >%s< value >%s<", identityKey, identityValue)
			}
		}
	}

	authzRoleMethods := authzRoleCache[hc.Method]
	if authzRoleMethods != nil {
		authzRoles := authzIdentityMethods[hc.Method]
		if authzRoles != nil {
			for _, roleName := range authzRoles {
				hasRole, err := rnr.hasContextRole(ctx, roleName)
				if err != nil {
					l.Warn("Failed checking context role >%s< >%v<", roleName, err)
					return err
				}
				if hasRole != true {
					msg := fmt.Sprintf("Context missing role >%s< value", roleName)
					l.Warn(msg)
					return fmt.Errorf(msg)
				}
				l.Info("Have role name >%s<", roleName)
			}
		}
	}

	return nil
}

// authzCacheConfig - cache authz configuration
func (rnr *Runner) authzCacheConfig(hc HandlerConfig) error {

	// Cache required identities
	if hc.MiddlewareConfig.AuthRequiredIdentities != nil {
		if authzIdentityCache == nil {
			authzIdentityCache = make(map[string]map[string][]string)
		}
		if authzIdentityCache[hc.Path] == nil {
			authzIdentityCache[hc.Path] = make(map[string][]string)
		}
		for _, authRequiredIdentity := range hc.MiddlewareConfig.AuthRequiredIdentities {
			authzIdentityCache[hc.Path][hc.Method] = append(authzIdentityCache[hc.Path][hc.Method], authRequiredIdentity)
		}
	}

	// Cache required roles
	if hc.MiddlewareConfig.AuthRequiredRoles != nil {
		if authzRoleCache == nil {
			authzRoleCache = make(map[string]map[string][]string)
		}
		if authzRoleCache[hc.Path] == nil {
			authzRoleCache[hc.Path] = make(map[string][]string)
		}
		for _, authRequiredRole := range hc.MiddlewareConfig.AuthRequiredRoles {
			authzRoleCache[hc.Path][hc.Method] = append(authzRoleCache[hc.Path][hc.Method], authRequiredRole)
		}
	}

	return nil
}
