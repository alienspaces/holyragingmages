package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/auth"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// auth
var authen *auth.Auth

// authenCache - path, method
var authenCache map[string][]string

// Authen -
func (rnr *Runner) Authen(path string, h HandlerFunc) (HandlerFunc, error) {

	var err error
	if authen == nil {
		authen, err = auth.NewAuth(rnr.Config, rnr.Log)
		if err != nil {
			rnr.Log.Warn("Failed new auth >%v<", err)
			return nil, err
		}
	}

	if authenCache == nil {
		for _, hc := range rnr.HandlerConfig {
			err := rnr.authenCacheConfig(hc)
			if err != nil {
				rnr.Log.Warn("Failed caching authen config >%v<", err)
				return nil, err
			}
		}
	}

	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Info("** Authen ** Checking authen types")

		// Authentication may add roles and identities to request context
		ctx := r.Context()

		// Authentication
		if authTypes, ok := authenCache[path]; ok {
			for _, authType := range authTypes {
				switch authType {
				case auth.AuthTypeJWT:
					l.Info("** Authen ** JWT")
					authString := r.Header.Get("Authorization")
					if authString == "" {
						msg := "Authorization header is empty"
						l.Warn(msg)
						rnr.WriteUnauthorizedError(l, w, fmt.Errorf(msg))
						return
					}
					if strings.Contains(authString, "Bearer ") {
						authString = strings.Split(authString, "Bearer ")[1]
					}
					claims, err := authen.DecodeJWT(authString)
					if err != nil {
						l.Warn("Failed authenticating token >%v<", err)
						rnr.WriteUnauthorizedError(l, w, err)
						return
					}

					l.Info("Have claims >%#v<", claims)

					ctx, err = rnr.addRolesContext(ctx, claims.Roles)
					if err != nil {
						l.Warn("Failed adding roles context >%v<", err)
						rnr.WriteUnauthorizedError(l, w, err)
						return
					}

					ctx, err = rnr.addIdentityContext(ctx, claims.Identity)
					if err != nil {
						l.Warn("Failed adding identity context >%v<", err)
						rnr.WriteUnauthorizedError(l, w, err)
						return
					}

				default:
					// Unsupported authentication configuration
					msg := "Unsupported authentication configuration"
					l.Warn(msg)
					rnr.WriteUnauthorizedError(l, w, fmt.Errorf(msg))
					return
				}
			}
		}

		h(w, r.WithContext(ctx), pp, qp, l, m)
	}

	return handle, nil
}

// authenCacheConfig - cache authen configuration
func (rnr *Runner) authenCacheConfig(hc HandlerConfig) error {

	if hc.MiddlewareConfig.AuthTypes != nil {

		if authenCache == nil {
			authenCache = make(map[string][]string)
		}
		if authenCache[hc.Path] == nil {
			authenCache[hc.Path] = []string{}
		}
		for _, authType := range hc.MiddlewareConfig.AuthTypes {
			authenCache[hc.Path] = append(authenCache[hc.Path], authType)
		}
	}

	return nil
}
