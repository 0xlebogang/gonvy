package httperr

import "net/http"

func Wrap(code, msg string, status int, cause error, details ...details) Error {
	return &appError{
		code:    code,
		status:  status,
		msg:     msg,
		details: details,
		cause:   cause,
	}
}

func BadRequest(cause error, msg string) Error {
	if cause == nil {
		panic("httperr: missing cause")
	}
	return Wrap(
		CodeBadRequest,
		writeMsg(CodeBadRequest, msg),
		http.StatusBadRequest,
		cause,
	)
}

func ValidationFailed(cause error, details ...details) Error {
	if cause == nil {
		panic("httperr: missing cause")
	}
	return Wrap(
		CodeValidationFailed,
		MsgByCode[CodeValidationFailed],
		http.StatusUnprocessableEntity,
		cause,
		details...,
	)
}

func Conflict(code, msg string) Error {
	return Wrap(
		code,
		writeMsg(code, msg),
		http.StatusConflict,
		nil,
	)
}

func Unauthorized(msg string) Error {
	return Wrap(
		CodeUnauthorized,
		writeMsg(CodeUnauthorized, msg),
		http.StatusUnauthorized,
		nil,
	)
}

func Forbidden() Error {
	return Wrap(
		CodeForbidden,
		MsgByCode[CodeForbidden],
		http.StatusForbidden,
		nil,
	)
}

func NotFound() Error {
	return Wrap(
		CodeNotFound,
		MsgByCode[CodeNotFound],
		http.StatusNotFound,
		nil,
	)
}

func MethodNotAllowed() Error {
	return Wrap(
		CodeMethodNotAllowed,
		MsgByCode[CodeMethodNotAllowed],
		http.StatusMethodNotAllowed,
		nil,
	)
}

func ServiceUnavailable(cause error) Error {
	if cause == nil {
		panic("httperr: missing cause")
	}
	return Wrap(
		CodeServiceUnavailable,
		MsgByCode[CodeServiceUnavailable],
		http.StatusServiceUnavailable,
		cause,
	)
}

func InternalServerError(cause error) Error {
	if cause == nil {
		panic("httperr: missing cause")
	}
	return Wrap(
		CodeInternalServerError,
		MsgByCode[CodeInternalServerError],
		http.StatusInternalServerError,
		cause,
	)
}

func writeMsg(code, msg string) string {
	genericMsg, ok := MsgByCode[code]
	if !ok {
		return msg
	}
	return genericMsg
}
