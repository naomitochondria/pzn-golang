package main

import (
	"fmt"
	"testing"
)

func random(input interface{}) interface{} {
	/*
		Process...
	*/
	return input
}

func TestTypeAssertion(t *testing.T) {

	// Mengubah interface kosong menjadi tipe data yang diinginkan
	var resultRandom interface{} = random("random string")
	var resultString string = resultRandom.(string)
	fmt.Println(resultString)

	// Menerima interface kosong
	result := random(false)
	switch value := result.(type) { // hanya bisa digunakan di switch
	case string:
		value += "!!"
		fmt.Println(value)
	case int:
		value *= 100
		fmt.Println(value)
	case bool:
		value = !value
		fmt.Println(value)
	default:
		fmt.Println("Unknown type")
	}
}
