package main

import (
	"fmt"
	"testing"
)

type Address struct {
	Country  string
	Province string
	City     string
}

func changeAddressToIndonesia(address *Address) *Address {
	address.Country = "Indondesia"

	return address
}

type BigStruct struct {
	Number1 uint64
	Number2 uint64
}

func (bigStruct *BigStruct) replaceNumber(number1, number2 uint64) {
	bigStruct.Number1 = bigStruct.Number1 * number1
	bigStruct.Number2 = bigStruct.Number2 * number2
}

func TestPointer(t *testing.T) {
	// reference by pointer
	var address1 Address = Address{
		Country:  "Indonesia",
		Province: "Jawa Tengah",
		City:     "Solo",
	}
	var address2 *Address = &address1
	var address3 Address = *address2 // Value structnya yang di pass
	address2.City = "Semarang"

	fmt.Println(address1, address2, address3)
	fmt.Println(address1, *address2, address3)

	// new keyword untuk membuat pointer
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	// Pointer di parameter function
	var address5 Address = Address{
		Country:  "Spain",
		Province: "Barcelona",
		City:     "Barcelona",
	}
	changeAddressToIndonesia(&address5)
	fmt.Println(address5)

	// Pointer di method struct
	veryBigStruct := BigStruct{
		1000000000, 900000000000,
	}
	veryBigStruct.replaceNumber(9000, 89999)
	fmt.Println(veryBigStruct)
}
