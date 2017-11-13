package errors

import "net/http"

// creates new instance of internal error
func NewInternalError(err error) *internalServerError {
	return &internalServerError{
		err: err,
	}
}

type internalServerError struct {
	err error
}

func (v *internalServerError) Error() string {
	return v.err.Error()
}

func (v *internalServerError) GetHTTPCode() int {
	return http.StatusInternalServerError
}

func (v *internalServerError) GetErrorType() string {
	return ErrorInternal
}

func (v *internalServerError) GetData(isDebug bool) interface{} {
	if isDebug {
		return map[string]interface{}{
			"msg": v.err.Error(),
			"err": v.err,
		}
	}

	return nil
}
