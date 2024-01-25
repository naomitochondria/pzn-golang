package main

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name          string
	DepartementID uint16
}

func processInput(i int) interface{} {
	if i > 1 && i < 10 {
		return 100
	} else if i >= 11 && i < 20 {
		return false
	} else if i >= 21 && i < 30 {
		return "21"
	} else {
		return Employee{"Employee A", 1009}
	}
}

func TestInterfaceKosong(t *testing.T) {
	// var getInput Employee = processInput(100)
	var getInput interface{} = processInput(100)
	fmt.Println(getInput)
}
