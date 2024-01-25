package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	map1 := map[int]string{
		100: "seratus",
		99:  "sembilan puluh sembilan",
		142: "seratus empat puluh dua",
	}
	map1[33] = "tiga puluh tiga"

	fmt.Println(map1)
	fmt.Println(len(map1))

	map2 := make(map[string]string)
	map2["nama"] = "tiga warna"
	delete(map2, "tiga warna")
	delete(map2, "nama")

	fmt.Println(map2)
}
