package error

type apiError string

const (
	TypeValidationError apiError = "validation_error"
)

type ApiError interface {
	GetErrorType() apiError
	GetHTTPCode() int
	Error() string
}