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
	if rnr.MiddlewareConfig.ValidateSchemaLocation == "" || rnr.MiddlewareConfig.ValidateSchemaMain == "" {
		rnr.Log.Info("Missing validate schema location or validate main schema, not validating request body")
		return h, nil
	}

	schemaLoc := rnr.MiddlewareConfig.ValidateSchemaLocation
	schema := rnr.MiddlewareConfig.ValidateSchemaMain
	schemaReferences := rnr.MiddlewareConfig.ValidateSchemaReferences

	rnr.Log.Info("Validating request body with schema %/%", schemaLoc, schema)

	// load and validate the schema
	sl := gojsonschema.NewSchemaLoader()
	sl.Validate = true

	// first load any referenced schemas
	for _, schemaName := range schemaReferences {
		rnr.Log.Info("Adding schema reference %/%", schemaLoc, schemaName)
		loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, schemaName))
		err := sl.AddSchemas(loader)
		if err != nil {
			rnr.Log.Warn("Failed adding schema reference >%v<", err)
			return nil, err
		}
	}

	// then load and compile the main schema (which references the other schemas)
	loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, schema))
	sd, err := sl.Compile(loader)
	if err != nil {
		rnr.Log.Warn("Failed compiling schema's >%v<", err)
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
			rnr.Log.Warn("Failed validate >%v<", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if result.Valid() != true {
			errStr := ""
			for _, e := range result.Errors() {
				errStr = fmt.Sprintf("%s, %s", errStr, e)
				rnr.Log.Warn("Invalid data >%s<", e)
			}
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		// delegate request
		h(w, r, ps)
	}

	return handle, nil
}
