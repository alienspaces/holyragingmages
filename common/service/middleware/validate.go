package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xeipuuv/gojsonschema"
)

// SchemaValidate -
func SchemaValidate(h httprouter.Handle, schemaLoc, mainSchema string, refSchemas ...string) (httprouter.Handle, error) {

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
		data := r.Context().Value(ContextDataKey)

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

	// schemaLoc := fmt.Sprintf("file:///%s/internal/services/schema/definitions", e.Get("APP_HOME"))

	// r := &Schema{
	// 	Env:              e,
	// 	Logger:           l.With().Str("package", PackageName).Logger(),
	// 	schemaDirectory:  schemaLoc,
	// 	schemaDefinition: sd,
	// }

	return handle, nil
}
