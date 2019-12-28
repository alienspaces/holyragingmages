package template

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// Request -
type Request struct {
}

// Response -
type Response struct {
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// MakeTemplateEndpoint -
func MakeTemplateEndpoint(svc Service) endpoint.Endpoint {

	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Request)
		resp, err := svc.Template(req)
		if err != nil {
			return Response{
				Err: err.Error(),
			}, nil
		}
		return resp, nil
	}
}

// DecodeRequest -
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeResponse -
func DecodeResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response Response
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// EncodeResponse -
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// EncodeRequest -
func EncodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
