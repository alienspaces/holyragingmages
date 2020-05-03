package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/spell/internal/record"
)

// GetSpellsHandler -
func (rnr *Runner) GetSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get spells handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Spell
	var err error

	id := p.ByName("id")

	// single resource
	if id != "" {

		rnr.Log.Info("Getting spell record ID >%s<", id)

		rec, err := m.(*model.Model).GetSpellRec(id, false)
		if err != nil {
			rnr.Log.Warn("Failed getting spell record >%v<", err)

			// model error
			res := rnr.ModelError(err)

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}

		// resource not found
		if rec == nil {
			rnr.Log.Warn("Get spell rec nil")

			// not found error
			res := rnr.NotFoundError(fmt.Errorf("Resource with ID >%s< not found", id))

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}

		recs = append(recs, rec)

	} else {

		rnr.Log.Info("Gatting all spell records")

		recs, err = m.(*model.Model).GetSpellRecs(nil, nil, false)
		if err != nil {

			// system error
			res := rnr.SystemError(err)

			err = rnr.WriteResponse(w, res)
			if err != nil {
				rnr.Log.Warn("Failed writing response >%v<", err)
				return
			}
			return
		}
	}

	// assign response properties
	data := []Data{}
	for _, rec := range recs {
		data = append(data, Data{
			ID:        rec.ID,
			CreatedAt: rec.CreatedAt,
			UpdatedAt: rec.UpdatedAt.Time,
		})
	}

	res := Response{
		Data: data,
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostSpellsHandler -
func (rnr *Runner) PostSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post spells handler ** p >%#v< m >#%v<", p, m)

	req := Request{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)

		// system error
		res := rnr.SystemError(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	rec := record.Spell{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).CreateSpellRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed creating spell record >%v<", err)

		// model error
		res := rnr.ModelError(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: []Data{
			{
				ID:        rec.ID,
				CreatedAt: rec.CreatedAt,
				UpdatedAt: rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutSpellsHandler -
func (rnr *Runner) PutSpellsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post spells handler ** p >%#v< m >#%v<", p, m)

	req := Request{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)

		// system error
		res := rnr.SystemError(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	rec := record.Spell{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).UpdateSpellRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed updating spell record >%v<", err)

		// model error
		res := rnr.ModelError(err)

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	// assign response properties
	res := Response{
		Data: []Data{
			{
				ID:        rec.ID,
				CreatedAt: rec.CreatedAt,
				UpdatedAt: rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}
