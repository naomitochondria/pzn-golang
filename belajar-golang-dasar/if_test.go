package main

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	name := "Tiga warna"
	if nameLen := len(name); nameLen != 0 {
		fmt.Println("Name is not empty!")
	}
}
