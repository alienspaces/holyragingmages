package payload

import (
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.com/alienspaces/holyragingmages/common/service"
	"gitlab.com/alienspaces/holyragingmages/common/type/logger"
	"gitlab.com/alienspaces/holyragingmages/common/type/payloader"
)

var _ payloader.Payloader = &Payload{}

// Payload -
type Payload struct {
	Log logger.Logger
}

// NewPayload -
func NewPayload(l logger.Logger) (*Payload, error) {

	p := Payload{
		Log: l,
	}

	return &p, nil
}

// ReadRequest -
func (p *Payload) ReadRequest(r *http.Request, s interface{}) error {

	data := r.Context().Value(service.ContextKeyData)

	if data != nil {
		r := strings.NewReader(data.(string))
		return json.NewDecoder(r).Decode(s)
	}

	return nil
}

// WriteResponse -
func (p *Payload) WriteResponse(w http.ResponseWriter, s interface{}) error {

	return json.NewEncoder(w).Encode(s)
}
