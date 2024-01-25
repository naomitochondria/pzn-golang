package test

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bags Bag[T]) {
	for _, bag := range bags {
		fmt.Println(bag)
	}
}

func TestGenericType(t *testing.T) {
	numbers := Bag[int]{1, 2, 3}
	PrintBag(numbers)

	letters := Bag[string]{"A", "B", "C"}
	PrintBag(letters)
}
