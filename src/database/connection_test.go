package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitCredits(t *testing.T) {
	// arrange
	credits := "driverName://user:pass@(tcp:3306)/dev"
	// act
	driver, dsn := splitCredits(credits)
	// assert
	assert.Equal(t, driver, "driverName")
	assert.Equal(t, dsn, "user:pass@(tcp:3306)/dev")
}
