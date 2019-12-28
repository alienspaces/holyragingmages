package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"gitlab.com/alienspaces/holyragingmages/service/template/internal/template"
)

func main() {

	listen := flag.String("port", "8080", "HTTP listen address")
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	// template
	var svc template.Service
	svc = template.ServiceHandler{}
	svc = template.LoggingMiddleware(logger)(svc)

	templateHandler := httptransport.NewServer(
		template.MakeTemplateEndpoint(svc),
		template.DecodeRequest,
		template.EncodeResponse,
	)

	http.Handle("/templates", templateHandler)

	logger.Log("info", fmt.Sprintf("Listening on http://localhost:%s", *listen))

	logger.Log("err", http.ListenAndServe(fmt.Sprintf(":%s", *listen), nil))
}
