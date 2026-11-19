package httperr

import "fmt"

type details struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

type Error interface {
	error
	GetCode() string
	GetStatus() int
	GetMsg() string
	GetDetails() []details
	AsResponse() *response
}
type appError struct {
	code    string
	status  int
	msg     string
	details []details
	cause   error
}

type response struct {
	Code    string    `json:"code"`
	Status  int       `json:"status"`
	Msg     string    `json:"message"`
	Details []details `json:"details"`
}

func (a *appError) AsResponse() *response {
	return &response{
		Code:    a.code,
		Status:  a.status,
		Msg:     a.msg,
		Details: a.details,
	}
}

func (a *appError) GetCode() string       { return a.code }
func (a *appError) GetStatus() int        { return a.status }
func (a *appError) GetMsg() string        { return a.msg }
func (a *appError) GetDetails() []details { return a.details }
func (a *appError) Unwrap() error         { return a.cause }

func (a *appError) Error() string {
	if a.cause == nil {
		return fmt.Sprintf("[%s]: %s", a.code, a.msg)
	}
	return fmt.Sprintf("[%s] %s: %s", a.code, a.msg, a.cause.Error())
}
