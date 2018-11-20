package primetrust

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimetrustVersion(t *testing.T) {
	assert.Equal(t, "1.0", Version, "Primetrust version does not match")
}
