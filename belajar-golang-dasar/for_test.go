package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	// Seperti while
	counter := 0

	for counter < 10 {
		fmt.Println(counter * 10)
		counter++
	}

	// Dengan init dan post statement
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// Range
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
	for ix, value := range months {
		fmt.Println(ix, "->", value)
	}
	for _, value := range months {
		fmt.Println(value)
	}

	map1 := map[int]string{
		100: "seratus",
		99:  "sembilan puluh sembilan",
		142: "seratus empat puluh dua",
	}
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	fmt.Println(map1)
}

var x = 1

func TestForV2(t *testing.T) {

	for i := &x; *i <= 10; *i++ {
		if *i == 4 {
			continue
		}

		if *i == 7 {
			break
		}

		fmt.Println(i)
	}

	fmt.Println(x)
}
