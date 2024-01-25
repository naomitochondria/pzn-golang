package main

import (
	"fmt"
	"testing"
)

func sum(numbers ...int) (sumValue int) {
	for _, n := range numbers {
		sumValue += n
	}

	return
}

func TestVariadicFunction(t *testing.T) {
	fmt.Println(sum(100, 200, 500))

	slice := []int{1000, 4000, 5000}
	fmt.Println(sum(slice...))
}
