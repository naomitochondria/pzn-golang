package main

import (
	"fmt"
	"testing"
)

func TestVariable(t *testing.T) {
	// Variable
	var numVar int16 = 89
	var floatVar float32 = 9.888

	var numVar2 = 89
	var floatVar2 = 3.034

	numVar3 := 392
	floatVar3 := 8.99

	var numVar4 uint8
	numVar4 = 88

	var (
		firstName = "Aba"
		lastName  = "Aca"
	)

	fmt.Println(numVar, floatVar, numVar2, floatVar2, numVar3, floatVar3, numVar4)
	fmt.Println(firstName, lastName)

	// Constant
	const PHI float32 = 3.4
	const GRAVITY = 9.8
	const (
		LEN_A = 10
		LEN_B = 19
	)

	// Datatype conversion
	var num64 int64 = 1000
	var num32 int32 = int32(num64)
	var num16 int16 = int16(num32)
	var num8 int8 = int8(num16)
	var num1 byte = 65
	var strNum1 string = string(num1)
	var str = "Abcdef"
	var str0 = string(str[0])

	fmt.Println(num64, num32, num16, num8, strNum1, str0)
}
