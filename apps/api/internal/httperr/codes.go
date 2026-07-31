package httperr

type ErrorCode string

const (
	CodeUnauthorized ErrorCode = "UNAUTHORIZED"
	CodeInvalidInput ErrorCode = "INVALID_INPUT"
	CodeEntityExists ErrorCode = "ENTITY_EXISTS"
	CodeNotFound     ErrorCode = "NOT_FOUND"
	CodeInternal     ErrorCode = "INTERNAL_ERROR"
)
