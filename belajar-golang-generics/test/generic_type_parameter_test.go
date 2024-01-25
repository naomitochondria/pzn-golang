package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// array/slice T dengan E yang merupakan apa saja
func Sum2[T []E, E interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | string
}](data T) E {
	var result E
	for _, d := range data {
		result += d
	}

	return result
}

func TestGenericTypeParameter(t *testing.T) {
	slice := []int{1, 2, 3}

	assert.Equal(t, 6, Sum2(slice))
}
