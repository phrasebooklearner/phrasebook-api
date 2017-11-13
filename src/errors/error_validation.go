package errors

import "net/http"

var _ ApiError = (*ValidationError)(nil)

func NewValidationError() *ValidationError {
	return &ValidationError{}
}

type ValidationError struct {
	errors []fieldError
}

type fieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// adds new filed validation error to an existing error list
func (v *ValidationError) AddFieldError(field, message string) *ValidationError {
	v.errors = append(v.errors, fieldError{
		Field: field,
		Error: message,
	})

	return v
}

// checks if there are no errors in the list
func (v *ValidationError) Empty() bool {
	return len(v.errors) == 0
}

func (v *ValidationError) Error() string {
	return ""
}

func (v *ValidationError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (v *ValidationError) GetErrorType() string {
	return ErrorValidation
}

func (v *ValidationError) GetData(isDebug bool) interface{} {
	return v.errors
}
