package response

import (
	"fmt"
	"net/http"
)

type InternalError struct {
	internalError error
}

var _ ApiError = (*InternalError)(nil)

func NewInternalError(err error) *InternalError {
	return &InternalError{
		internalError: err,
	}
}

func (v InternalError) Error() string {
	return fmt.Sprintf("Something went wrong")
}

func (v InternalError) GetHTTPCode() int {
	return http.StatusBadRequest
}

func (v InternalError) GetErrorType() string {
	return "internal_error"
}

func (v InternalError) GetData(isDebug bool) interface{} {

	data := make(map[string]string)

	if isDebug {
		data["error"] = v.internalError.Error()
	}

	return data

}
