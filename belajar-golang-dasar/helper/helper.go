package helper

import "fmt"

func init() {
	fmt.Println("Calling helper package...")
}

func Multiply(a, b int) int {
	return a * b
}
