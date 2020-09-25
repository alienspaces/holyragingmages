package runner

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"

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

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString("abcdefgh")
	if err != nil {
		l.Warn("Failed singing JWT >%v<", err)
		rnr.WriteSystemError(l, w, nil)
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
