package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	int | int32 | int64 | float32 | float64
}](first T, second T) T {
	if first <= second {
		return first
	}

	return second
}

func TestInlineConstraint(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, 0.8, FindMin(0.8025, 0.8))
}
