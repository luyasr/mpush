package errs

import (
	"errors"
	"fmt"
)

const (
	defaultErrorCode = 1
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func New(code int, format string, a ...any) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

func GetCode(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}

	return defaultErrorCode
}

func GetMessage(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Error()
	}

	return err.Error()
}
