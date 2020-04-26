package server

import "fmt"

// ErrorCode -
const (
	ErrorCodeSystem       string = "internal_error"
	ErrorDetailSystem     string = "An internal error has occurred"
	ErrorCodeValidation   string = "validation_error"
	ErrorDetailValidation string = "Request contains validation errors"
	ErrorCodeNotFound     string = "not_found"
	ErrorDetailNotFound   string = "Requested resource could not be found"
)

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
