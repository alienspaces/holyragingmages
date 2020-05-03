package runner

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/holyragingmages/common/payload"
	"gitlab.com/alienspaces/holyragingmages/common/server"
	"gitlab.com/alienspaces/holyragingmages/common/type/modeller"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
	"gitlab.com/alienspaces/holyragingmages/common/type/runnable"
	"gitlab.com/alienspaces/holyragingmages/service/item/internal/model"
	"gitlab.com/alienspaces/holyragingmages/service/item/internal/record"
)

// Runner -
type Runner struct {
	server.Runner
}

// Response -
type Response struct {
	server.Response
	Data []Data `json:"data"`
}

// Request -
type Request struct {
	server.Request
	Data Data `json:"data"`
}

// Data -
type Data struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Fault -
type Fault struct {
	Error error
}

// ensure we comply with the Runnerer interface
var _ runnable.Runnable = &Runner{}

// NewRunner -
func NewRunner() *Runner {

	r := Runner{}

	r.RouterFunc = r.Router
	r.MiddlewareFunc = r.Middleware
	r.HandlerFunc = r.Handler
	r.ModellerFunc = r.Modeller
	r.PayloaderFunc = r.Payloader

	r.HandlerConfig = []server.HandlerConfig{
		{
			Method:           http.MethodGet,
			Path:             "/api/items",
			HandlerFunc:      r.GetItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
		},
		{
			Method:           http.MethodGet,
			Path:             "/api/items/:id",
			HandlerFunc:      r.GetItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{},
		},
		{
			Method:      http.MethodPost,
			Path:        "/api/items",
			HandlerFunc: r.PostItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
		{
			Method:      http.MethodPut,
			Path:        "/api/items/:id",
			HandlerFunc: r.PutItemsHandler,
			MiddlewareConfig: server.MiddlewareConfig{
				ValidateSchemaLocation: "schema",
				ValidateSchemaMain:     "main.schema.json",
				ValidateSchemaReferences: []string{
					"data.schema.json",
				},
			},
		},
	}

	return &r
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

// Modeller -
func (rnr *Runner) Modeller() (modeller.Modeller, error) {

	rnr.Log.Info("** Item Model **")

	m, err := model.NewModel(rnr.Config, rnr.Log, rnr.Store)
	if err != nil {
		rnr.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// Payloader -
func (rnr *Runner) Payloader() (payloader.Payloader, error) {

	rnr.Log.Info("** Payloader **")

	p, err := payload.NewPayload()
	if err != nil {
		rnr.Log.Warn("Failed new payloader >%v<", err)
		return nil, err
	}

	return p, nil
}

// Handler - default handler
func (rnr *Runner) Handler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Item handler **")

	fmt.Fprint(w, "Hello from item!\n")
}

// GetItemsHandler -
func (rnr *Runner) GetItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Get items handler ** p >%#v< m >%#v<", p, m)

	var recs []*record.Item
	var err error

	id := p.ByName("id")

	// single resource
	if id != "" {

		rnr.Log.Info("Getting item record ID >%s<", id)

		rec, err := m.(*model.Model).GetItemRec(id, false)
		if err != nil {
			rnr.Log.Warn("Failed getting item record >%v<", err)

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
			rnr.Log.Warn("Get item rec nil")

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

		rnr.Log.Info("Gatting all item records")

		recs, err = m.(*model.Model).GetItemRecs(nil, nil, false)
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

// PostItemsHandler -
func (rnr *Runner) PostItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post items handler ** p >%#v< m >#%v<", p, m)

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

	rec := record.Item{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).CreateItemRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed creating item record >%v<", err)

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

// PutItemsHandler -
func (rnr *Runner) PutItemsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params, m modeller.Modeller) {

	rnr.Log.Info("** Post items handler ** p >%#v< m >#%v<", p, m)

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

	rec := record.Item{}

	// assign request properties
	rec.ID = req.Data.ID

	err = m.(*model.Model).UpdateItemRec(&rec)
	if err != nil {
		rnr.Log.Warn("Failed updating item record >%v<", err)

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
