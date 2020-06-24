package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/common/payload"
	"gitlab.com/alienspaces/holyragingmages/server/common/server"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/payloader"
)

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, l logger.Logger, m modeller.Modeller) {

	l.Info("** Spell handler **")

	fmt.Fprint(w, "Hello from spell!\n")
}

// Router -
func (rnr *Runner) Router(r *httprouter.Router) error {

	rnr.Log.Info("** Spell Router **")

	return nil
}

// Middleware -
func (rnr *Runner) Middleware(h server.Handle) (server.Handle, error) {

	rnr.Log.Info("** Spell Middleware **")

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
