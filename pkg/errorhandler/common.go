package errorhandler

func NewAuthFailed(format string, a ...any) *Error {
	return New(401, format, a...)
}

func NewExists(format string, a ...any) *Error {
	return New(400, format, a...)
}

func NewUpdateFailed(format string, a ...any) *Error {
	return New(400, format, a...)
}

func NewNotFound(format string, a ...any) *Error {
	return New(404, format, a...)
}
