package errors

import "net/http"

// creates new instance of http error
func NewHTTPError(code int) *httpError {
	return &httpError{
		code:    code,
	}
}

type httpError struct {
	code    int
}

func (v httpError) GetHTTPCode() int {
	return v.code
}

func (v httpError) Error() string {
	return http.StatusText(v.code)
}

func (v httpError) GetErrorType() string {
	return ErrorHTTP
}

func (v httpError) GetData(isDebug bool) interface{} {
	return map[string]interface{} {
		"code": v.code,
	}
}
