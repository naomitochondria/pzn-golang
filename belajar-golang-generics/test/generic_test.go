package test

import (
	"fmt"
	"testing"
)

/*
	Contoh kode non-generic
	Logicnya sama tapi tipe datanya beda,
	sehingga harus membuat dua fungsi yang berbeda
*/

func SumInt(slice []int) int {
	var sum int = 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func SumFloat(slice []float64) float64 {
	var sum float64 = 0
	for _, s := range slice {
		sum += s
	}
	return sum
}

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestGenerics(t *testing.T) {
	fmt.Println(SumInt([]int{1, 2, 3}))
	fmt.Println(SumFloat([]float64{1.5, 2.5, 3.5}))

	fmt.Println(Length[int](9))
}
