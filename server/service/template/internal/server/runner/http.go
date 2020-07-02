package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/core/payload"
	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/payloader"
)

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Template handler **")

	fmt.Fprint(w, "Hello from template!\n")
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Template Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h server.Handle) (server.Handle, error) {

	rnr.Log.Info("** Template Middleware **")

	return h, nil
}

// Payloader -
func (rnr *Runner) Payloader(l logger.Logger) (payloader.Payloader, error) {

	l.Info("** Payloader **")

	p, err := payload.NewPayload()
	if err != nil {
		l.Warn("Failed new payloader >%v<", err)
		return nil, err
	}

	return p, nil
}
