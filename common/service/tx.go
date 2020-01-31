package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Tx -
func (rnr *Runner) Tx(h Handle) (Handle, error) {

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, _ Modeller) {

		rnr.Log.Warn("** Tx ** beginning transaction")

		tx, err := rnr.Store.GetTx()
		if err != nil {
			rnr.Log.Warn("Failed getting DB connection >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		if rnr.ModelFunc == nil {
			rnr.Log.Warn("Runner NewModelFunc is nil")
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		m, err := rnr.ModelFunc(rnr.Config, rnr.Log, rnr.Store)
		if err != nil {
			rnr.Log.Warn("Failed new model >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		err = m.Init(tx)
		if err != nil {
			rnr.Log.Warn("Failed init model >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		// delegate request
		h(w, r, ps, m)

		rnr.Log.Warn("** Tx ** committing transaction")

		// TODO: determine how to decide whether to commit or rollback..
		err = tx.Commit()
		if err != nil {
			rnr.Log.Warn("Failed Tx commit >%v<", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
	}

	return handle, nil
}
