package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeipuuv/gojsonschema"

	"gitlab.com/alienspaces/holyragingmages/server/core/type/logger"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/modeller"
)

// schemaCache - path, method, schema
var schemaCache map[string]map[string]*gojsonschema.Schema

// queryParamCache - path, method, []string
var queryParamCache map[string]map[string][]string

// Validate -
func (rnr *Runner) Validate(path string, h HandlerFunc) (HandlerFunc, error) {

	rnr.Log.Info("** Validate ** cache query param lists")

	// cache query parameter lists
	if queryParamCache == nil {
		for _, hc := range rnr.HandlerConfig {
			err := rnr.cacheQueryParamList(hc)
			if err != nil {
				rnr.Log.Warn("Failed caching query param list >%v<", err)
				return nil, err
			}
		}
	}

	rnr.Log.Info("** Validate ** cache schemas")

	// load configured schemas
	if schemaCache == nil {
		for _, hc := range rnr.HandlerConfig {
			err := rnr.validateCacheSchemas(hc)
			if err != nil {
				rnr.Log.Warn("Failed loading schemas >%v<", err)
				return nil, err
			}
		}
	}

	handle := func(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

		l.Debug("** Validate ** request URI >%s< method >%s<", r.RequestURI, r.Method)

		// validate query parameters
		qp, err := rnr.validateQueryParameters(l, path, r)
		if err != nil {
			rnr.WriteResponse(l, w, rnr.ValidationError(err))
			return
		}

		// only validate requests where the method provides data
		if r.Method != http.MethodPost && r.Method != http.MethodPut {
			l.Debug("Skipping validation of URI >%s< method >%s<", r.RequestURI, r.Method)

			// delegate request
			h(w, r, pp, qp, l, m)
			return
		}

		// schema for URI and method
		s := schemaCache[path][r.Method]
		if s == nil {
			l.Debug("Not validating URI >%s< method >%s<", r.RequestURI, r.Method)

			// delegate request
			h(w, r, pp, qp, l, m)
			return
		}

		// data from context
		data := r.Context().Value(ContextKeyData)

		l.Debug("Data >%s<", data)

		// load the data
		var dataLoader gojsonschema.JSONLoader
		switch data.(type) {
		case nil:
			l.Warn("Data is nil")
			rnr.WriteResponse(l, w, rnr.SystemError(fmt.Errorf("Data is nil")))
			return
		case string:
			dataLoader = gojsonschema.NewStringLoader(data.(string))
		default:
			dataLoader = gojsonschema.NewGoLoader(data)
		}

		// validate the data
		result, err := s.Validate(dataLoader)
		if err != nil {
			l.Warn("Failed validate >%v<", err)
			if err.Error() == "EOF" {
				rnr.WriteResponse(l, w, rnr.ValidationError(fmt.Errorf("Posted data is empty, check request content")))
				return
			}
			rnr.WriteResponse(l, w, rnr.SystemError(err))
			return
		}

		if result.Valid() != true {
			errStr := ""
			for _, e := range result.Errors() {
				l.Warn("Invalid data >%s<", e)
				if errStr == "" {
					errStr = e.String()
					continue
				}
				errStr = fmt.Sprintf("%s, %s", errStr, e.String())
			}
			rnr.WriteResponse(l, w, rnr.ValidationError(fmt.Errorf("%s", errStr)))
			return
		}

		// delegate request
		h(w, r, pp, qp, l, m)
	}

	return handle, nil
}

// validateQueryParameters - validate any provided parameters
func (rnr *Runner) validateQueryParameters(l logger.Logger, path string, r *http.Request) (map[string]interface{}, error) {

	if len(queryParamCache) == 0 {
		l.Debug("Handler method >%s< path >%s< not configured with params list", r.Method, path)
		return nil, nil
	}

	// query parameters
	q := r.URL.Query()

	// valid query parameters
	qp := map[string]interface{}{}

	// allowed parameters
	params := queryParamCache[path][r.Method]

	for paramName, paramValue := range q {
		l.Info("Checking parameter >%s< >%s<", paramName, paramValue)

		found := false
		for _, param := range params {
			if paramName == param {
				found = true
			}
		}
		if found != true {
			msg := fmt.Sprintf("Parameter >%s< not allowed", paramName)
			l.Warn(msg)
			return nil, fmt.Errorf(msg)
		}
		qp[paramName] = paramValue
	}

	l.Info("All parameters okay")

	return qp, nil
}

// cacheQueryParamList -
func (rnr *Runner) cacheQueryParamList(hc HandlerConfig) error {

	if len(hc.MiddlewareConfig.ValidateQueryParams) == 0 {
		rnr.Log.Info("Handler method >%s< path >%s< not configured with params list", hc.Method, hc.Path)
		return nil
	}

	rnr.Log.Info("Handler method >%s< path >%s< has params list >%#v<", hc.Method, hc.Path, hc.MiddlewareConfig.ValidateQueryParams)

	if queryParamCache == nil {
		queryParamCache = map[string]map[string][]string{}
	}
	if queryParamCache[hc.Path] == nil {
		queryParamCache[hc.Path] = make(map[string][]string)
	}
	queryParamCache[hc.Path][hc.Method] = hc.MiddlewareConfig.ValidateQueryParams

	return nil
}

// validateCacheSchemas - load validation JSON schemas
func (rnr *Runner) validateCacheSchemas(hc HandlerConfig) error {

	if hc.MiddlewareConfig.ValidateSchemaLocation == "" || hc.MiddlewareConfig.ValidateSchemaMain == "" {
		rnr.Log.Info("Handler method >%s< path >%s< not configured for validation", hc.Method, hc.Path)
		return nil
	}

	schemaLoc := hc.MiddlewareConfig.ValidateSchemaLocation
	schema := hc.MiddlewareConfig.ValidateSchemaMain
	schemaReferences := hc.MiddlewareConfig.ValidateSchemaReferences

	schemaPath := rnr.Config.Get("APP_SERVER_SCHEMA_PATH")
	schemaLoc = fmt.Sprintf("file://%s/%s", schemaPath, schemaLoc)

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
