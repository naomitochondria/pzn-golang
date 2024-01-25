package main

import (
	"fmt"
	"strings"
	"testing"
)

type Filter func(string) string

func sayHello(name string, filter Filter) {
	filteredName := filter(name)

	fmt.Println("Halo,", filteredName)
}

func filterBadWords(name string) string {
	if strings.Contains(strings.ToLower(name), "anjing") {
		name = "[SENSOR]"
	}

	return name
}

func TestFunctionAsParameter(t *testing.T) {
	sayHello("Pembunuh anjing", filterBadWords)
	sayHello("Cupa", func(name string) string {
		return name + "!!"
	})
}
