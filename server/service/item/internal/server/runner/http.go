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
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Item handler **")

	fmt.Fprint(w, "Hello from item!\n")
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Item Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h server.Handle) (server.Handle, error) {

	rnr.Log.Info("** Item Middleware **")

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
