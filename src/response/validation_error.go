package response

import (
	"fmt"
	"net/http"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

var _ ApiError = (*ValidationError)(nil)

func NewValidationError(field string, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
	}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Error in field [%s]: %s", v.Field, v.Message)
}

func (v ValidationError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (v ValidationError) GetErrorType() apiError {
	return TypeValidationError
}
