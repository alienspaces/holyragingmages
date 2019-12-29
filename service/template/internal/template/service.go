package template

import (
	"errors"
)

// Service provides operations on strings.
type Service interface {
	Template(req Request) (Response, error)
}

// ServiceHandler -
type ServiceHandler struct{}

// Template -
func (ServiceHandler) Template(req Request) (Response, error) {

	// TODO: implement all support HTTP methods
	return Response{
		Test: req.Test,
	}, nil
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// ServiceMiddleware is a chainable behaviour modifier for TemplateService.
type ServiceMiddleware func(Service) Service
