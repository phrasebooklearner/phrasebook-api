package response

type ApiError interface {
	GetErrorType() string
	GetHTTPCode() int
	Error() string
	GetData(isDebug bool) interface{}
}
