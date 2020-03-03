package runner

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/template/internal/record"
)

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Template handler **")

	fmt.Fprint(w, "Hello from template!\n")
}

// GetTemplatesHandler -
func (rnr *Runner) GetTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get templates handler ** p >%#v< m >%#v<", p, m)

	fmt.Fprint(w, "Hello from GET templates handler!\n", p)
}

// PostTemplatesHandler -
func (rnr *Runner) PostTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post templates handler ** p >%#v< m >#%v<", p, m)

	req := Request{}

	err := rnr.ReadRequest(r, &req)
	if err != nil {
		rnr.Log.Warn("Failed reading request >%v<", err)
		fmt.Fprint(w, "Failed reading request\n", err)
		return
	}

	rec := record.Template{}
	rec.ID = req.Data.ID

	err = m.(*model.Model).CreateTemplateRec(&rec)
	if err != nil {
		res := Response{
			Error: err,
		}
		err = rnr.WriteResponse(w, &res)
		if err != nil {
			rnr.Log.Warn("Failed writing response >%v<", err)
			fmt.Fprint(w, "Failed writing response\n", err)
			return
		}
		return
	}

	res := Response{
		Data: Data{
			ID: req.Data.ID,
		},
	}

	err = rnr.WriteResponse(w, &res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		fmt.Fprint(w, "Failed writing response\n", err)
		return
	}
}

// PutTemplatesHandler -
func (rnr *Runner) PutTemplatesHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	data := r.Context().Value(service.ContextKeyData)

	rnr.Log.Info("** Put templates handler ** p >%#v< m >#%v< data >%v<", p, m, data)

	fmt.Fprint(w, "Hello from Put templates handler!\n", p, "\n", data)
}
