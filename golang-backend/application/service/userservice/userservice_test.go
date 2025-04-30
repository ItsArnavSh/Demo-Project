package userservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparePassword(t *testing.T) {
	hash, _ := HashPassword("mypassword")
	err := ComparePassword("mypassword", hash)
	assert.NoError(t, err, "Password should match")
}
