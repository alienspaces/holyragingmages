package payloader

import "net/http"

// Payloader -
type Payloader interface {
	ReadRequest(r *http.Request, s interface{}) error
	WriteResponse(w http.ResponseWriter, s interface{}) error
}
