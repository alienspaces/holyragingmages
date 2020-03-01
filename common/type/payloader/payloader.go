package payloader

// Payloader -
type Payloader interface {
	Request(jsonData string) (structData interface{}, err error)
	Response(structData interface{}) (jsonData string, err error)
}
