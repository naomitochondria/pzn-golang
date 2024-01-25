package test

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](params1 T1, params2 T2) (T1, T2) {
	fmt.Println(params1)
	fmt.Println(params2)

	return params1, params2
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[int, string](9000, "Abc")
	MultipleParameter[float64, bool](90.09, false)
}
