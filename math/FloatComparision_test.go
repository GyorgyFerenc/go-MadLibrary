package Math_test

import (
	"testing"

	m "github.com/GyorgyFerenc/go-MadLibrary/math"

	"github.com/stretchr/testify/assert"
)

func TestFloatComparision(t *testing.T) {
	assert := assert.New(t)

	assert.True(m.CompareFloats(0, 0))

	a := 5.0
	b := a - m.Threshold
	assert.True(m.CompareFloats(a, b))

	b = a - m.Threshold - 0.0000000001
	assert.False(m.CompareFloats(a, b))
}
