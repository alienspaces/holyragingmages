package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/server/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/common/type/modeller"
)

// Tx -
func (rnr *Runner) Tx(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, l logger.Logger, _ modeller.Modeller) {

		l.Info("** Tx ** beginning database transaction")

		tx, err := rnr.Store.GetTx()
		if err != nil {
			l.Warn("Failed getting DB connection >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		l.Warn("Have DB tx >%#v<", tx)

		// NOTE: The PREPARER is created an initialised with every request instead of
		// creating and assigning to a runner struct "Prepare" property at start up.
		// This ensures statements are valid for the current database transaction.

		// preparer
		if rnr.PreparerFunc == nil {
			l.Warn("Runner PreparerFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		p, err := rnr.PreparerFunc(l)
		if err != nil {
			l.Warn("Failed PreparerFunc >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if p == nil {
			l.Warn("Preparer is nil, cannot continue")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = p.Init(tx)
		if err != nil {
			l.Warn("Failed init preparer >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// NOTE: The modeller is created an initialised with every request instead of
		// creating and assigning to a runner struct "Model" property at start up.
		// This prevents directly accessing a shared property from with the handler
		// function which is running in a goroutine. Otherwise accessing the "Model"
		// property would require locking and block simultaneous requests.

		// modeller
		if rnr.ModellerFunc == nil {
			l.Warn("Runner ModellerFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		m, err := rnr.ModellerFunc(l)
		if err != nil {
			l.Warn("Failed ModellerFunc >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if m == nil {
			l.Warn("Modeller is nil, cannot continue")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = m.Init(p, tx)
		if err != nil {
			l.Warn("Failed init modeller >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// delegate request
		h(w, r, ps, l, m)

		l.Info("** Tx ** committing database transaction")

		// TODO: Handle should return a possible error so we can
		// determine whether we need to commit or rollback
		err = tx.Commit()
		if err != nil {
			l.Warn("Failed Tx commit >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
	}

	return handle, nil
}
