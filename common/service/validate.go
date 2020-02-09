package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeipuuv/gojsonschema"

	"gitlab.com/alienspaces/holyragingmages/common/modeller"
)

// schemaCache - path, method, schema
var schemaCache map[string]map[string]*gojsonschema.Schema

// Validate -
func (rnr *Runner) Validate(path string, h Handle) (Handle, error) {

	rnr.Log.Info("** Validate ** loading schemas")

	// load configured schemas
	if schemaCache == nil {
		for _, hc := range rnr.HandlerConfig {
			err := rnr.validateLoadSchemas(hc)
			if err != nil {
				rnr.Log.Warn("Failed loading schemas >%v<", err)
				return nil, err
			}
		}
	}

	handle := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, m modeller.Modeller) {

		rnr.Log.Info("** Validate ** request URI >%s< method >%s<", r.RequestURI, r.Method)

		// schema for URI and method
		s := schemaCache[path][r.Method]
		if s == nil {
			rnr.Log.Info("Not validating URI >%s< method >%s<", r.RequestURI, r.Method)

			// delegate request
			h(w, r, ps, m)
			return
		}

		// NOTE: may want to skip methods that don't support body data at all

		// data from context
		data := r.Context().Value(ContextKeyData)

		rnr.Log.Info("Data >%s<", data)

		// load the data
		var dataLoader gojsonschema.JSONLoader
		switch data.(type) {
		case nil:
			http.Error(w, "Data is nil", http.StatusBadRequest)
			return
		case string:
			dataLoader = gojsonschema.NewStringLoader(data.(string))
		default:
			dataLoader = gojsonschema.NewGoLoader(data)
		}

		// validate the data
		result, err := s.Validate(dataLoader)
		if err != nil {
			rnr.Log.Warn("Failed validate >%v<", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if result.Valid() != true {
			errStr := ""
			for _, e := range result.Errors() {
				rnr.Log.Warn("Invalid data >%s<", e)
				if errStr == "" {
					errStr = e.String()
					continue
				}
				errStr = fmt.Sprintf("%s, %s", errStr, e.String())
			}
			http.Error(w, errStr, http.StatusBadRequest)
			return
		}

		// delegate request
		h(w, r, ps, m)
	}

	return handle, nil
}

func (rnr *Runner) validateLoadSchemas(hc HandlerConfig) error {

	if hc.MiddlewareConfig.ValidateSchemaLocation == "" || hc.MiddlewareConfig.ValidateSchemaMain == "" {
		rnr.Log.Info("Handler method >%s< path >%s< not configured for validation", hc.Method, hc.Path)
		return nil
	}

	schemaLoc := hc.MiddlewareConfig.ValidateSchemaLocation
	schema := hc.MiddlewareConfig.ValidateSchemaMain
	schemaReferences := hc.MiddlewareConfig.ValidateSchemaReferences

	appHome := rnr.Config.Get("APP_HOME")
	schemaLoc = fmt.Sprintf("file://%s/%s", appHome, schemaLoc)

	rnr.Log.Info("Loading schema %s/%s", schemaLoc, schema)

	// load and validate the schema
	sl := gojsonschema.NewSchemaLoader()
	sl.Validate = true

	// first load any referenced schemas
	for _, schemaName := range schemaReferences {
		rnr.Log.Info("Adding schema reference %s/%s", schemaLoc, schemaName)
		loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, schemaName))
		err := sl.AddSchemas(loader)
		if err != nil {
			rnr.Log.Warn("Failed adding schema reference %v", err)
			return err
		}
	}

	// then load and compile the main schema (which references the other schemas)
	loader := gojsonschema.NewReferenceLoader(fmt.Sprintf("%s/%s", schemaLoc, schema))
	s, err := sl.Compile(loader)
	if err != nil {
		rnr.Log.Warn("Failed compiling schema's >%v<", err)
		return err
	}

	if schemaCache == nil {
		schemaCache = map[string]map[string]*gojsonschema.Schema{}
	}
	if schemaCache[hc.Path] == nil {
		schemaCache[hc.Path] = make(map[string]*gojsonschema.Schema)
	}
	schemaCache[hc.Path][hc.Method] = s

	return nil
}
