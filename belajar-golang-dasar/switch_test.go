package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	// Dengan kondisi
	name := "Tiga warna"

	switch nameLen := len(name); nameLen < 10 {
	case true:
		fmt.Println("Nama terlalu pendek")
	case false:
		fmt.Println("Nama pas")
	}

	// Tanpa kondisi
	switch {
	case name == "Putih":
		fmt.Println("Dia gendut!")
	case name == "Cassiopeia":
		fmt.Println("Dia lucu banget!")
	case name == "Cupa":
		fmt.Println("Dia cantik!")
	case name == "Tiga warna":
		fmt.Println("Dia ibuk!")
	}
}
