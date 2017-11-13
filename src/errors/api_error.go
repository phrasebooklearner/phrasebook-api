package errors

type ApiError interface {
	Error() string
	GetErrorType() string
	GetHTTPCode() int
	GetData(isDebug bool) interface{}
}

const (
	ErrorValidation string = "validation-error"
	ErrorInternal   string = "internal-error"
	ErrorHTTP       string = "http-error"
)
