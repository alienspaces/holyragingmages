package server

import (
	"fmt"
	"net/http"
)

// ErrorCode -
const (
	ErrorCodeSystem       string = "internal_error"
	ErrorDetailSystem     string = "An internal error has occurred"
	ErrorCodeValidation   string = "validation_error"
	ErrorDetailValidation string = "Request contains validation errors"
	ErrorCodeNotFound     string = "not_found"
	ErrorDetailNotFound   string = "Requested resource could not be found"
)

// WriteModelError -
func (rnr *Runner) WriteModelError(w http.ResponseWriter, err error) {

	rnr.Log.Warn("Model error >%v<", err)

	// model error
	res := rnr.ModelError(err)

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
	return
}

// ModelError -
func (rnr *Runner) ModelError(err error) Response {

	rnr.Log.Error("Error >%v<", err)

	return Response{
		Error: ResponseError{
			Code:   ErrorCodeValidation,
			Detail: err.Error(),
		},
	}
}

// WriteSystemError -
func (rnr *Runner) WriteSystemError(w http.ResponseWriter, err error) {

	rnr.Log.Warn("System error >%v<", err)

	// system error
	res := rnr.SystemError(err)

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
	return
}

// SystemError -
func (rnr *Runner) SystemError(err error) Response {

	rnr.Log.Error("Error >%v<", err)

	// NOTE: never expose actual system error details
	return Response{
		Error: ResponseError{
			Code:   ErrorCodeSystem,
			Detail: ErrorDetailSystem,
		},
	}
}

// ValidationError -
func (rnr *Runner) ValidationError(err error) Response {

	rnr.Log.Error("Error >%v<", err)

	if err == nil {
		err = fmt.Errorf(ErrorDetailValidation)
	}

	return Response{
		Error: ResponseError{
			Code:   ErrorCodeValidation,
			Detail: err.Error(),
		},
	}
}

// WriteNotFoundError -
func (rnr *Runner) WriteNotFoundError(w http.ResponseWriter, id string) {

	err := fmt.Errorf("Resource with ID >%s< not found", id)

	rnr.Log.Warn("Not found error >%v<", err)

	// not found error
	res := rnr.NotFoundError(err)

	err = rnr.WriteResponse(w, res)
	if err != nil {
		rnr.Log.Warn("Failed writing response >%v<", err)
		return
	}
	return
}

// NotFoundError -
func (rnr *Runner) NotFoundError(err error) Response {

	rnr.Log.Error("Error >%v<", err)

	if err == nil {
		err = fmt.Errorf(ErrorDetailNotFound)
	}

	return Response{
		Error: ResponseError{
			Code:   ErrorCodeNotFound,
			Detail: err.Error(),
		},
	}
}
