package httperr

const (
	CodeBadRequest          = "BAD_REQUEST"
	CodeValidationFailed    = "VALIDATION_FAILED"
	CodeConflict            = "CONFLICT"
	CodeUnauthorized        = "UNAUTHORIZED"
	CodeForbidden           = "FORBIDDEN"
	CodeNotFound            = "NOT_FOUND"
	CodeMethodNotAllowed    = "METHOD_NOT_ALLOWED"
	CodeServiceUnavailable  = "SERVICE_UNAVAILABLE"
	CodeInternalServerError = "INTERNAL_SERVER_ERROR"
)

var MsgByCode = map[string]string{
	CodeBadRequest:          "The provided input is malformed. Check your input and try again.",
	CodeValidationFailed:    "One or more fields failed validation. Check your input and try again.",
	CodeConflict:            "The provided input conflicts with an existing resource. Check your input and try again.",
	CodeUnauthorized:        "A valid session is required to perform this action. Log in and try again.",
	CodeForbidden:           "You do not have permission to perform this action.",
	CodeNotFound:            "The resource you requested could not be found.",
	CodeMethodNotAllowed:    "Method not allowed for this resource.",
	CodeServiceUnavailable:  "The service is temporarily unavailable. Please try again later.",
	CodeInternalServerError: "An unexpected error has occurred. Please try again later.",
}
