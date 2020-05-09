package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
)

// Tx -
func (rnr *Runner) Tx(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, p httprouter.Params, _ modeller.Modeller) {

		rnr.Log.Info("** Tx ** beginning database transaction")

		tx, err := rnr.Store.GetTx()
		if err != nil {
			rnr.Log.Warn("Failed getting DB connection >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = rnr.Prepare.Init(tx)
		if err != nil {
			rnr.Log.Warn("Failed init preparer >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// NOTE: The model is created an initialised with every request instead of
		// creating and assigning to a runner struct "Model" property at start up.
		// This prevents directly accessing a shared property from with the handler
		// function which is running in a goroutine. Otherwise accessing the "Model"
		// property would require locking and block simultaneous requests.

		// modeller
		if rnr.ModellerFunc == nil {
			rnr.Log.Warn("Runner ModellerFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		m, err := rnr.ModellerFunc()
		if err != nil {
			rnr.Log.Warn("Failed ModellerFunc >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = m.Init(rnr.Prepare, tx)
		if err != nil {
			rnr.Log.Warn("Failed init modeller >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// delegate request
		h(w, r, p, m)

		rnr.Log.Info("** Tx ** committing database transaction")

		// TODO: Handle should return a possible error so we can
		// determine whether we need to commit or rollback
		err = tx.Commit()
		if err != nil {
			rnr.Log.Warn("Failed Tx commit >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
	}

	return handle, nil
}