package response

import "phrasebook-api/src/errors"

func Success() *apiResponse {
	return &apiResponse{
		Success: true,
	}
}

func ApiError(error errors.ApiError, debug bool) *apiResponse {
	return &apiResponse{
		Error: &apiError{
			Type: error.GetErrorType(),
			Text: error.Error(),
			Data: error.GetData(debug),
		},
	}
}

type apiResponse struct {
	Success bool      `json:"success"`
	Error   *apiError `json:"error,omitempty"`
}

type apiError struct {
	Type string      `json:"type,omitempty"`
	Text string      `json:"text,omitempty"`
	Data interface{} `json:"data,omitempty"`
}
