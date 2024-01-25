package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// capacity masih cukup, maka akan append->replace element di array yang di refer
	var slice1 = months[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = append(slice1, "BUKAN MEI")
	fmt.Println(slice2)
	fmt.Println(months)

	// capacity tidak cukup, akan merefer ke array baru
	var slice4 = months[11:]
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	var slice5 = append(slice4, "KIAMAT")
	fmt.Println(slice5)
	fmt.Println(months)

	// copy slice by value
	array1 := []int16{
		1, 3, 5, 7, 9,
	}
	array2 := make([]int16, len(array1), cap(array1))
	copy(array2, array1)
	fmt.Println(array1, array2)

	// slice vs array
	iniArray1 := [...]int{
		9, 7, 5,
	}
	iniArray2 := [3]int{
		3, 5, 7,
	}
	iniSlice := []int{
		2, 4, 6,
	}

	fmt.Println(iniArray1, iniArray2, iniSlice)
}
