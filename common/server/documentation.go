package server

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GenerateHandlerDocumentation - generates documentationbased on handler configuration
func (rnr *Runner) GenerateHandlerDocumentation() ([]byte, error) {

	rnr.Log.Info("** Generate Handler Documentation **")

	var b strings.Builder

	fmt.Fprintf(&b, `
<html>
<head>
<style>
	body {
		background-color: #efefef;
	}
	h2 {
		color: #504ecc;
	}
	pre {
		background-color: #ffffff;
		padding: 10px;
	}
	.path-method {
		color: #629153;
	}
</style>
</head>
<body>
	`)

	fmt.Fprintf(&b, "<h2>Documentation</h2>")

	for _, config := range rnr.HandlerConfig {

		if config.DocumentationConfig.Document != true {
			// skip documenting this endpoint
			continue
		}

		var schemaMainContent []byte
		var schemaDataContent []byte
		var err error

		appHome := rnr.Config.Get("APP_SERVER_HOME")
		schemaLoc := config.MiddlewareConfig.ValidateSchemaLocation
		if schemaLoc != "" {

			schemaMain := config.MiddlewareConfig.ValidateSchemaMain
			filename := fmt.Sprintf("%s/%s/%s", appHome, schemaLoc, schemaMain)

			rnr.Log.Info("Schema main content filename >%s<", filename)

			schemaMainContent, err = ioutil.ReadFile(filename)
			if err != nil {
				return nil, err
			}

			schemaReferences := config.MiddlewareConfig.ValidateSchemaReferences
			for _, schemaReference := range schemaReferences {

				filename := fmt.Sprintf("%s/%s/%s", appHome, schemaLoc, schemaReference)

				rnr.Log.Info("Schema reference content filename >%s<", filename)

				schemaDataContent, err = ioutil.ReadFile(filename)
				if err != nil {
					return nil, err
				}
			}
		}

		fmt.Fprintf(&b, "<div class='path'><h4><span class='path-method'>%s</span> - <span class='path=url'>%s</span></h4></div>", config.Method, config.Path)
		if config.DocumentationConfig.Description != "" {
			fmt.Fprintf(&b, "<div class='description'>%s</div>", config.DocumentationConfig.Description)
		}
		if schemaMainContent != nil {
			fmt.Fprintf(&b, "<div class='schema'><h4>Schema</h4></div><pre class='schema-data'>%s</pre>", string(schemaMainContent))
		}
		if schemaDataContent != nil {
			fmt.Fprintf(&b, "<pre class='schema-data'>%s</pre>", string(schemaDataContent))
		}
	}

	fmt.Fprintf(&b, `
	</body>
		`)

	return []byte(b.String()), nil
}
