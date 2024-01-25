package main

import (
	"fmt"
	"testing"
)

func EndAppSafely() {
	message := recover()
	fmt.Println(message)
	fmt.Println("Successfully end app...")
}

func RunApp(isError bool) {
	defer EndAppSafely() // Recover harus ada di defer

	fmt.Println("Starting app...")
	if isError {
		panic("ERROR")
	}
	fmt.Println("Running app...")
}

func TestPanicDeferRecover(t *testing.T) {
	RunApp(true)
	fmt.Println("Test")
}
