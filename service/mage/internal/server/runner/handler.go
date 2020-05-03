package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/mage/internal/record"
)

// GetMagesHandler -
func (rnr *Runner) GetMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get mages handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Mage
	var err error

	id := p.ByName("id")

	// single resource
	if id != "" {

		rnr.Log.Info("Fetching resource ID >%s<", id)

		rec, err := m.(*model.Model).GetMageRec(id, false)
		if err != nil {
			rnr.Log.Warn("Failed getting mage record >%v<", err)

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
			rnr.Log.Warn("Get mage rec nil")

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

		rnr.Log.Info("Getting all template records")

		recs, err = m.(*model.Model).GetMageRecs(nil, nil, false)
		if err != nil {
			rnr.Log.Warn("Failed getting mage records >%v<", err)

			// model error
			res := rnr.ModelError(err)

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
			ID:           rec.ID,
			Name:         rec.Name,
			Strength:     rec.Strength,
			Dexterity:    rec.Dexterity,
			Intelligence: rec.Intelligence,
			Experience:   rec.Experience,
			Coin:         rec.Coin,
			CreatedAt:    rec.CreatedAt,
			UpdatedAt:    rec.UpdatedAt.Time,
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

// PostMagesHandler -
func (rnr *Runner) PostMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post mages handler ** p >%#v< m >#%v<", p, m)

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

	rec := record.Mage{}

	// assign request properties
	rec.ID = req.Data.ID
	rec.Name = req.Data.Name

	err = m.(*model.Model).CreateMageRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed creating mage record >%v<", err)

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
				ID:           rec.ID,
				Name:         rec.Name,
				Strength:     rec.Strength,
				Dexterity:    rec.Dexterity,
				Intelligence: rec.Intelligence,
				Experience:   rec.Experience,
				Coin:         rec.Coin,
				CreatedAt:    rec.CreatedAt,
				UpdatedAt:    rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutMagesHandler -
func (rnr *Runner) PutMagesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Put mages handler ** p >%#v< m >#%v<", p, m)

	id := p.ByName("id")

	rnr.Log.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetMageRec(id, false)
	if err != nil {
		rnr.Log.Warn("Failed getting mage record >%v<", err)

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
		rnr.Log.Warn("Get mage rec nil")

		// not found error
		res := rnr.NotFoundError(fmt.Errorf("Resource with ID >%s< not found", id))

		err = rnr.WriteResponse(w, res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			return
		}
		return
	}

	req := Request{}

	err = rnr.ReadRequest(r, &req)
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

	// assign request properties
	rec.Name = req.Data.Name

	err = m.(*model.Model).UpdateMageRec(rec)
	if err != nil {
		rnr.Log.Warn("Failed updating mage record >%v<", err)

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
				ID:           rec.ID,
				Name:         rec.Name,
				Strength:     rec.Strength,
				Dexterity:    rec.Dexterity,
				Intelligence: rec.Intelligence,
				Experience:   rec.Experience,
				Coin:         rec.Coin,
				CreatedAt:    rec.CreatedAt,
				UpdatedAt:    rec.UpdatedAt.Time,
			},
		},
	}

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
}
