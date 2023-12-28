package shared

import (
	"fmt"
	"net/http"
	"strings"
)

type ResponseError interface {
	error
	HttpStatus() int
}

type GenericError struct {
	ErrorMsg       string
	HttpStatusCode int
}

func (e GenericError) Error() string {
	return strings.TrimSpace(e.ErrorMsg)
}

func (e GenericError) HttpStatus() int {
	return e.HttpStatusCode
}

func MapError(exitCode int, stderr string) ResponseError {
	switch exitCode {
	case 11:
		return GenericError{HttpStatusCode: http.StatusNotFound, ErrorMsg: stderr}
	default:
		return GenericError{HttpStatusCode: http.StatusInternalServerError, ErrorMsg: fmt.Sprintf("Exit Code %d: %s", exitCode, stderr)}
	}
}
