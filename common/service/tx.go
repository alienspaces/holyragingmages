package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/modeller"
)

// Tx -
func (rnr *Runner) Tx(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, _ modeller.Modeller) {

		rnr.Log.Warn("** Tx ** beginning transaction")

		tx, err := rnr.Store.GetTx()
		if err != nil {
			rnr.Log.Warn("Failed getting DB connection >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if rnr.PreparerFunc == nil {
			rnr.Log.Warn("Runner PreparerFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		p, err := rnr.PreparerFunc(rnr.Log, tx)
		if err != nil {
			rnr.Log.Warn("Failed PreparerFunc >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if rnr.ModelFunc == nil {
			rnr.Log.Warn("Runner ModelFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		m, err := rnr.ModelFunc(rnr.Config, rnr.Log, rnr.Store)
		if err != nil {
			rnr.Log.Warn("Failed ModelFunc >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = m.Init(p, tx)
		if err != nil {
			rnr.Log.Warn("Failed init model >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// delegate request
		h(w, r, ps, m)

		rnr.Log.Warn("** Tx ** committing transaction")

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
