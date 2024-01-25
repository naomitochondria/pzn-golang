package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

type Age int

func IsSmaller[T Number](number1, number2 T) bool {
	if number1 < number2 {
		return true
	}

	return false
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, true, IsSmaller[int](9, 100))
	assert.Equal(t, true, IsSmaller[float64](float64(9.045), float64(100.1025)))
	assert.Equal(t, true, IsSmaller[Age](100, 90000))

	// type inference
	assert.Equal(t, true, IsSmaller(9, 100))
	assert.Equal(t, true, IsSmaller(float64(9.045), float64(100.1025)))
	assert.Equal(t, true, IsSmaller(100, 90000))
}
