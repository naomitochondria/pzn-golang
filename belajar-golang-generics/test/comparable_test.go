package test

import (
	"fmt"
	"testing"
)

// // Any tidak bisa di komparasikan
// func IsSame[T any](param1 T, param2 T) bool {
// 	if param1 == param2 {

// 	}
// }

func IsSame[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	}

	return false
}

func TestComparable(t *testing.T) {
	fmt.Println(IsSame[int](9, 10000))
}
