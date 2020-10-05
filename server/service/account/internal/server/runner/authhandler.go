package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/constant"
	"gitlab.com/alienspaces/holyragingmages/server/core/auth"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/schema"
	"gitlab.com/alienspaces/holyragingmages/server/service/account/internal/model"
)

// PostAuthHandler -
func (rnr *Runner) PostAuthHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post auth handler ** p >%#v< m >#%v<", pp, m)

	req := schema.AuthRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	accountRec, err := m.(*model.Model).VerifyProviderToken(model.AuthData{
		Provider:          req.Data.Provider,
		ProviderAccountID: req.Data.ProviderAccountID,
		ProviderToken:     req.Data.ProviderToken,
		AccountEmail:      req.Data.AccountEmail,
		AccountName:       req.Data.AccountName,
	})
	if err != nil {
		rnr.WriteUnauthorizedError(l, w, err)
		return
	}

	a, err := auth.NewAuth(rnr.Config, rnr.Log)
	if err != nil {
		rnr.WriteUnauthorizedError(l, w, err)
		return
	}

	// TODO: Expand on account roles
	roles := []string{
		constant.AuthRoleDefault,
	}

	identity := map[string]interface{}{
		constant.AuthIdentityAccount: accountRec.ID,
	}

	claims := auth.Claims{
		Roles:    roles,
		Identity: identity,
	}

	tokenString, err := a.EncodeJWT(&claims)
	if err != nil {
		rnr.WriteUnauthorizedError(l, w, err)
		return
	}

	// assign response properties
	res := schema.AuthResponse{
		Data: []schema.AuthData{
			{
				Provider:          req.Data.Provider,
				ProviderAccountID: req.Data.ProviderAccountID,
				ProviderToken:     req.Data.ProviderToken,
				AccountID:         accountRec.ID,
				AccountName:       accountRec.Name,
				AccountEmail:      accountRec.Email,
				Token:             tokenString,
			},
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}
