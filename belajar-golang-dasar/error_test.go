package main

import (
	"errors"
	"fmt"
	"testing"
)

/**
Asal dari sini:
type error interface {
	Error() string
}
*/

func division(a, b int) (float32, error) {
	if b == 0 {
		return 0, errors.New("Division by zero!")
	} else {
		return float32(a) / float32(b), nil
	}
}

func TestError(t *testing.T) {
	var zeroDivision error = errors.New("Division by zero!")
	fmt.Println(zeroDivision)

	if result, err := division(11, 2); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("ERROR:", err)
	}
}
