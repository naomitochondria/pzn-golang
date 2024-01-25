package main

import "testing"

// Tidak bisa diakses dari luar package
var vesion = "1.0.0"

// Bisa diakses dari luar package
var Application = "Golang"

type Config struct {
	Name      string // Bisa dipanggil config.Name
	DateAdded string // Bisa dipanggil config.DateAdded
	key       string // panic
}

func TestAccessModifier(t *testing.T) {
	/*
		Pass
	*/
}
