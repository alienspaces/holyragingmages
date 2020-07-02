package payload

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/alienspaces/holyragingmages/server/core/server"
	"gitlab.com/alienspaces/holyragingmages/server/core/type/payloader"
)

var _ payloader.Payloader = &Payload{}

// Payload -
type Payload struct{}

// NewPayload -
func NewPayload() (*Payload, error) {

	p := Payload{}

	return &p, nil
}

// ReadRequest -
func (p *Payload) ReadRequest(r *http.Request, s interface{}) error {

	data := r.Context().Value(server.ContextKeyData)

	if data != nil {
		r := strings.NewReader(data.(string))
		err := json.NewDecoder(r).Decode(s)
		if err != nil {
			// Include data in error response
			return fmt.Errorf("Failed decoding request data >%s< >%v<", data.(string), err)
		}
	}

	return nil
}

// WriteResponse -
func (p *Payload) WriteResponse(w http.ResponseWriter, status int, s interface{}) error {

	// content type json
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// status
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(s)
}
