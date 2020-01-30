package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeipuuv/gojsonschema"
)

// Validate -
func (rnr *Runner) Validate(h httprouter.Handle) (httprouter.Handle, error) {

	// JSON schema validatation
	if rnr.MiddlewareConfig.ValidateSchemaLocation == "" || rnr.MiddlewareConfig.ValidateMainSchema == "" {
		rnr.Log.Info("Missing validate schema location or validate main schema, not validating request body")
		return h, nil
	}

	schemaLoc := rnr.MiddlewareConfig.ValidateSchemaLocation
	mainSchema := rnr.MiddlewareConfig.ValidateMainSchema
	refSchemas := rnr.MiddlewareConfig.ValidateReferenceSchemas

	rnr.Log.Info("Validating request body with schema %/%", schemaLoc, mainSchema)

	// load and validate the schema
	sl := gojsonschema.NewSchemaLoader()
	sl.Validate = true

	// first load any referenced schemas
	for _, schemaName := range refSchemas {
		loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, schemaName))
		err := sl.AddSchemas(loader)
		if err != nil {
			return nil, err
		}
	}

	// then load and compile the main schema (which references the other schemas)
	loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, mainSchema))
	sd, err := sl.Compile(loader)
	if err != nil {
		return nil, err
	}

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// data from context
		data := r.Context().Value(ContextKeyData)

		// load the data
		var dataLoader gojsonschema.JSONLoader
		switch data.(type) {
		case nil:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case string:
			dataLoader = gojsonschema.NewStringLoader(data.(string))
		default:
			dataLoader = gojsonschema.NewGoLoader(data)
		}

		// validate the data
		result, err := sd.Validate(dataLoader)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if result.Valid() != true {
			errStr := ""
			for _, e := range result.Errors() {
				errStr = fmt.Sprintf("%s, %s", errStr, e)
			}
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		// delegate request
		h(w, r, ps)
	}

	return handle, nil
}
