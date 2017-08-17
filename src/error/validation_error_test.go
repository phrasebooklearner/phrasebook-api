package error

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationError_GetHTTPCode(t *testing.T) {
	// arrange
	err := NewValidationError("field", "error")
	// assert
	assert.Equal(t, http.StatusOK, err.GetHTTPCode())
}

func TestValidationError_GetErrorType(t *testing.T) {
	// arrange
	//err := NewValidationError("field", "error")
	// assert
	assert.Equal(t, TypeValidationError, "fail test!")
}
